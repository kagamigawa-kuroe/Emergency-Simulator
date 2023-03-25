package rooms

import (
	"log"
	"math/rand"
	"strconv"
	"sync"
	"time"

	"gitlab.utc.fr/wanhongz/emergency-simulator/agent/patient"
)

/// 挂号室
type ReceptionRoom struct {
	sync.Mutex
	QueueNumber        int              				// Nombre de files d'attente activées
	Queues             map[string]chan *patient.Patient // toutes les files d'attente
	QueuesLength       map[string]int                   // La longueur de chaque file d'attente
	QueuesDoctor       map[string]*ReceptionDoctor      // Médecins par cohorte
	DocorsQueue        map[*ReceptionDoctor]string      // File d'attente de mappage des médecins
	AllPatientsWaiting map[string][]*patient.Patient    // tous les patients en attente
	MsgRequest         chan *patient.Patient            // canal de demande
	MsgDoctor          chan *ReceptionDoctor            // Canal de rétroaction des médecins
}

func (rr* ReceptionRoom) GetQueuesNumber() int{
	return len(rr.Queues)
}

func (rr* ReceptionRoom) AddQueue(){
	rr.Lock()
	rr.QueueNumber++
	i := rr.QueueNumber
	rr.QueuesDoctor["Queue"+strconv.FormatInt(int64(i), 10)].Lock()
	rr.QueuesDoctor["Queue"+strconv.FormatInt(int64(i), 10)].WorkOrNot = true
	rr.QueuesDoctor["Queue"+strconv.FormatInt(int64(i), 10)].Unlock()
	rr.Unlock()
	log.Println("a new queue waiting start")
}

func (rr* ReceptionRoom) ReduceQueue() {
	rr.Lock()
	i := rr.QueueNumber
	rr.QueueNumber--
	rr.QueuesDoctor["Queue"+strconv.FormatInt(int64(i), 10)].Lock()
	rr.QueuesDoctor["Queue"+strconv.FormatInt(int64(i), 10)].WorkOrNot = false
	rr.QueuesDoctor["Queue"+strconv.FormatInt(int64(i), 10)].Unlock()
	rr.Unlock()
	log.Println("a new queue waiting stop")
}

/// médecin pour reception
type ReceptionDoctor struct {
	sync.Mutex
	WorkOrNot        bool                  // s'il faut travailler
	ID               int                   // identifiant unique
	status           int                   // 1 occupé 0 libre
	QueueResponsable chan *patient.Patient // file d'attente responsable
	Msgreturn        chan *ReceptionDoctor // canal de rétroaction
}

func ( r *ReceptionDoctor ) Getstatus() int {
	return r.status
}

func (rr *ReceptionRoom) HandlerRquest(p *patient.Patient) {
	// Trouvez la position la plus appropriée et placez-la
	rr.Lock()
	m := rr.QueuesLength["Queue1"]
	qq := "Queue1"
	for i, j := range rr.QueuesLength {
		rr.QueuesDoctor[i].Lock()
		if j < m && rr.QueuesDoctor[i].WorkOrNot == true {
			m = j
			qq = i
		}
		rr.QueuesDoctor[i].Unlock()
	}

	rr.QueuesLength[qq]++
	rr.Queues[qq] <- p
	rr.AllPatientsWaiting[qq] = append(rr.AllPatientsWaiting[qq], p)
	log.Println("Patient" + strconv.FormatInt(int64(p.ID), 10) + " join the " + qq)
	rr.Unlock()
}

func (rr *ReceptionRoom) Run() {
	log.Println("ReceptionRoom start working")
	for _, j := range rr.QueuesDoctor {
		go j.Run()
	}
	for {
		select {
		case n := <-rr.MsgRequest:
			rr.HandlerRquest(n)
		case m := <-rr.MsgDoctor:
			rr.HandlerDoctor(m)
		default:
			time.Sleep(1 * time.Second)
		}
	}
}

func (rr *ReceptionRoom) HandlerDoctor(r *ReceptionDoctor) {
	rr.Lock()

	qq := rr.DocorsQueue[r]

	// fmt.Println(len(rr.AllPatientsWaiting[qq]))

	rr.QueuesLength[qq]--
	if len(rr.AllPatientsWaiting[qq]) != 0 {
		rr.AllPatientsWaiting[qq] = rr.AllPatientsWaiting[qq][1:]
	}

	rr.Unlock()
}

func (rd *ReceptionDoctor) HandlerPatientRequest(patient2 *patient.Patient) {
	rd.status = 1
	log.Println("ReceptionDoctor" + strconv.FormatInt(int64(rd.ID), 10) + " start dealing with patient " + strconv.FormatInt(int64(patient2.ID), 10))

	// Simulez le temps d'enregistrement, ajoutez aléatoire
	time.Sleep(time.Duration(rand.Int31n(3)+3) * time.Second)
	patient2.Msg_receive_reception <- "ticket"

	rd.status = 0
	rd.Msgreturn <- rd
}

func (rd *ReceptionDoctor) Run() {
	log.Println("ReceptionDoctor" + strconv.FormatInt(int64(rd.ID), 10) + " start working")
	for {
		select {
		case n, ok := <-rd.QueueResponsable:
			if !ok {
				log.Println("ReceptionDoctor" + strconv.FormatInt(int64(rd.ID), 10) + " stop")
				return
			}
			// log.Println("ReceptionDoctor" + strconv.FormatInt(int64(rd.ID), 10) + " start dealing with patient " + strconv.FormatInt(int64(n.ID), 10))
			rd.HandlerPatientRequest(n)
		default:
			time.Sleep(1 * time.Second)
		}
	}
}

var (
	instance *ReceptionRoom
	once     sync.Once
)

func GetInstance(n int) *ReceptionRoom {
	once.Do(func() {
		instance = &ReceptionRoom{
			QueueNumber:        3,
			Queues:             make(map[string]chan *patient.Patient),
			QueuesLength:       make(map[string]int),
			QueuesDoctor:       make(map[string]*ReceptionDoctor),
			DocorsQueue:        make(map[*ReceptionDoctor]string),
			AllPatientsWaiting: make(map[string][]*patient.Patient),
			MsgRequest:         make(chan *patient.Patient, 20),
			MsgDoctor:          make(chan *ReceptionDoctor, 20),
		}
		for i := 1; i <= n; i++ {
			c := make(chan *patient.Patient, 20)
			instance.Queues["Queue"+strconv.FormatInt(int64(i), 10)] = c
			instance.QueuesLength["Queue"+strconv.FormatInt(int64(i), 10)] = 0
			p := &ReceptionDoctor{
				WorkOrNot:        true,
				ID:               i,
				status:           0,
				QueueResponsable: c,
				Msgreturn:        instance.MsgDoctor,
			}
			if i > 3 {
				p.WorkOrNot = false
			}
			instance.DocorsQueue[p] = "Queue" + strconv.FormatInt(int64(i), 10)
			instance.QueuesDoctor["Queue"+strconv.FormatInt(int64(i), 10)] = p
			instance.AllPatientsWaiting["Queue"+strconv.FormatInt(int64(i), 10)] = make([]*patient.Patient, 0)
		}
	})
	return instance
}
