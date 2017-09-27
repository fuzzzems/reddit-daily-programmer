package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	messages := make(map[int]*Message)
	packetChan := make(chan Packet)
	msgChan := make(chan *Message)
	done := make(chan bool)

	go func() {
		for scanner.Scan() {
			if scanner.Text() == "" {
				close(packetChan)
				done <- true
				return
			}
			p, err := toPacket(scanner.Text())
			if err != nil {
				log.Fatal(err)
			}
			packetChan <- p
		}
	}()

	for {
		select {
		case p := <-packetChan:
			if p.messageID > 0 {
				msg := addPacketToMessage(p, messages[p.messageID])
				messages[p.messageID] = msg
				if msg.maxSize == msg.curSize {
					go func(msg *Message) {
						msgChan <- msg
					}(msg)
				}
			}
		case msg := <-msgChan:
			fmt.Println(msg.toString())
		case <-done:
			close(msgChan)
			close(done)
			return
		}
	}
}

func toPacket(line string) (Packet, error) {
	var p Packet

	var messageID, packetID, maxSize int

	_, err := fmt.Sscanf(line, "%d %d %d", &messageID, &packetID, &maxSize)
	if err != nil {
		return p, err
	}

	p.messageID, p.packetID, p.maxSize, p.data = messageID, packetID, maxSize, line

	return p, nil
}

type Packet struct {
	messageID int
	packetID  int
	maxSize   int
	data      string
}

type Message struct {
	maxSize int
	packets []Packet
	id      int
	curSize int
}

func (m Message) toString() string {
	msg := ""
	fmt.Println()
	for _, v := range m.packets {
		msg += fmt.Sprintf("%s \n", v.data)
	}
	return msg
}

func addPacketToMessage(p Packet, msg *Message) *Message {
	if msg == nil {
		msg = &Message{
			packets: make([]Packet, p.maxSize+1),
			maxSize: p.maxSize,
			curSize: 0,
			id:      p.messageID,
		}
	}
	msg.packets[p.packetID] = p
	msg.curSize++
	return msg
}
