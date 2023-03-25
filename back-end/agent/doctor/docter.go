package doctor

import (
	"log"
	"sync"
	"time"
)

type Doctor struct {
	sync.Mutex                  // mutex
	ID          int             // ID unique du médecin
	Specialized map[string]bool // spécialité du médecin
	Usable      bool            // Est-il actuellement disponible
	Ability     int             // Valeur de la capacité du médecin 1-5 points
}

type DoctorManager struct {
	sync.Mutex
	AllDoctor      []*Doctor
	DoctorReuqest  chan int     // Le type int envoyé représente la gravité du patient
	DoctorResponce chan *Doctor // retourner les médecins disponibles
}

func (dm *DoctorManager) AddDoctor(level int) {
	dm.Lock()

	d := NewDoctor(len(dm.AllDoctor)+1, make(map[string]bool), true, level)
	dm.AllDoctor = append(dm.AllDoctor, d)

	log.Println("Add a new doctor")
	dm.Unlock()
}

func (dm *DoctorManager) DeleteDoctor(level int) {
	dm.Lock()
	flag := false
	for !flag {
		for i, j := range dm.AllDoctor {
			j.Lock()
			if j.Usable == true && j.Ability == level {
				// 删除j
				if i == 0 {
					dm.AllDoctor = dm.AllDoctor[1:]
				} else if i == len(dm.AllDoctor) {
					dm.AllDoctor = dm.AllDoctor[:len(dm.AllDoctor)-1]
				} else {
					dm.AllDoctor = append(dm.AllDoctor[:i], dm.AllDoctor[i+1:]...)
				}

				flag = true
				j.Unlock()
				break
			}
			j.Unlock()
		}
	}

	log.Println("Delete a new doctor")
	dm.Unlock()
}

func (dm *DoctorManager) handlerRequest(value int) {
	for i := 0; i < len(dm.AllDoctor); i++ {
		// Déterminer s'il y a des médecins satisfaits à leur tour
		if dm.AllDoctor[i].Usable == true && dm.AllDoctor[i].Ability >= value {
			dm.DoctorResponce <- dm.AllDoctor[i]
			return
		}
	}
	dm.DoctorResponce <- nil
	return
}

var (
	instance *DoctorManager
	once     sync.Once
)

func GetDoctorManagerInstance(n int) *DoctorManager {
	once.Do(func() {
		instance = &DoctorManager{
			AllDoctor:      make([]*Doctor, 0),
			DoctorReuqest:  make(chan int, 20),
			DoctorResponce: make(chan *Doctor, 20),
		}
		for i := 1; i <= n; i++ {
			d := NewDoctor(i, make(map[string]bool), true, i)
			instance.AllDoctor = append(instance.AllDoctor, d)
		}
	})
	return instance
}

func (dm *DoctorManager) Run() {
	log.Println("DoctorManager start working")
	for {
		select {
		case i := <-dm.DoctorReuqest:
			dm.handlerRequest(i)
		default:
			time.Sleep(1 * time.Second)
		}
	}
}

// Constructeur
func NewDoctor(id int, specialized map[string]bool, usable bool, ability int) *Doctor {
	return &Doctor{
		ID:          id,
		Specialized: specialized,
		Usable:      true,
		Ability:     ability,
	}
}

// définit l'état disponible
func (d *Doctor) SetUsable(s bool) {
	d.Lock()
	d.Usable = s
	d.Unlock()
}
