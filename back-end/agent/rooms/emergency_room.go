package rooms

import (
	"log"
	"strconv"
	"sync"
	"time"
)

type EmergencyRoom struct {
	sync.Mutex
	Level  int
	Status int // 0 libre 1 occupé
	ID     int // identifiant unique
}

// salle d'attente détermine les besoins du patient pendant la traversée
// Envoyer une demande à l'agent de gestion des urgences pour déterminer s'il existe un type de chambre correspondant
type EmergencyRoomManager struct {
	sync.Mutex
	WorkNumber int                       // nombre de pièces sur lesquelles travailler
	RoomList   map[string]*EmergencyRoom // toutes les chambres
	MsgRequest chan int                  // canal de réception des demandes de chambre int représente le niveau de la chambre souhaitée
	MsgReponse chan *EmergencyRoom       // Le canal de réponse utilisé pour renvoyer les chambres disponibles ne renvoie pas nil
}

func (erm *EmergencyRoomManager) AddRoom(level int) {
	erm.Lock()
	erm.WorkNumber++
	i := erm.WorkNumber
	erm.RoomList["EmergencyRoom"+strconv.FormatInt(int64(i), 10)].Level = level
	log.Println("Start a new room, now there are totally " + strconv.FormatInt(int64(erm.WorkNumber), 10) + " work")
	for i:=1; i <= erm.WorkNumber; i++ {
		log.Println(erm.RoomList["EmergencyRoom"+strconv.FormatInt(int64(i), 10)].Level)
	}
	log.Println("------")
	erm.Unlock()
}

func (erm *EmergencyRoomManager) DeleteRoom() {
	erm.Lock()
	i := erm.WorkNumber
	erm.WorkNumber--
	erm.RoomList["EmergencyRoom"+strconv.FormatInt(int64(i), 10)].Level = -1
	log.Println("Stop a new room, now there are totally " + strconv.FormatInt(int64(erm.WorkNumber), 10) + " work")
	erm.Unlock()
}

func (erm *EmergencyRoomManager) Run() {
	log.Println("EmergencyRoomManager start working")
	for {
		select {
		case i := <-erm.MsgRequest:
			log.Println("EmergencyRoomCenter get a new request of level " + strconv.FormatInt(int64(i), 10))
			// 不能用协程
			erm.check(i)
		default:
			time.Sleep(1 * time.Second)
		}
	}
}

var (
	instance_erm *EmergencyRoomManager
	once2        sync.Once
)

func GetEmergencyRoomManagerInstance(n int) *EmergencyRoomManager {
	once2.Do(func() {
		instance_erm = &EmergencyRoomManager{
			WorkNumber: 5,
			RoomList:   make(map[string]*EmergencyRoom),
			MsgRequest: make(chan int, 20),
			MsgReponse: make(chan *EmergencyRoom, 20),
		}

		for i := 1; i <= 5; i++ {
			instance_erm.RoomList["EmergencyRoom"+strconv.FormatInt(int64(i), 10)] = &EmergencyRoom{
				Level:  i,
				Status: 0,
				ID:     i,
			}
			log.Println("EmergencyRoom" + strconv.FormatInt(int64(i), 10) + " has been create")
		}

		for i := 6; i <= n; i++ {
			instance_erm.RoomList["EmergencyRoom"+strconv.FormatInt(int64(i), 10)] = &EmergencyRoom{
				Level:  -1,
				Status: 0,
				ID:     i,
			}
			log.Println("EmergencyRoom" + strconv.FormatInt(int64(i), 10) + " has been create")
		}
	})
	return instance_erm
}

func (erm *EmergencyRoomManager) check(LevelNeed int) {
	erm.Lock()
	var ans *EmergencyRoom = nil

	// Vérifier le niveau de chaque pièce à tour de rôle
	for _, v := range erm.RoomList {
		if v.Status == 0 && v.Level != -1 {
			if ans == nil && v.Level >= LevelNeed {
				ans = v
			} else if ans != nil && v.Level >= LevelNeed && ans.Level > v.Level {
				ans = v
			}

		}
	}

	erm.MsgReponse <- ans
	erm.Unlock()
}
