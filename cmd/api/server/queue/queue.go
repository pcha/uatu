package queue

import (
	"log"

	"the-one/internal/saver"
)

type Queue struct {
	channel chan Queueable
	saver   saver.Saver
	stop    chan bool
}

type Queueable struct {
	Bucket string
	Data   *saver.Fact
}

func New(s saver.Saver) *Queue {
	return &Queue{
		make(chan Queueable),
		s,
		make(chan bool),
	}
}

func (q *Queue) StartListening() {
	go q.listen()
}

func (q *Queue) StopListening() {
	q.stop <- true
}

func (q *Queue) listen() {
	for {
		select {
		case o := <-q.channel:
			err := q.saver.Save(o.Data, o.Bucket)
			if err != nil {
				log.Fatal(err)
			}
		case <-q.stop:
			break
		}
	}
}

func (q *Queue) Receive(bucket string, data *saver.Fact) {
	q.channel <- Queueable{
		bucket,
		data,
	}
}

func (q *Queue) GetSaver() saver.Saver {
	return q.saver
}
