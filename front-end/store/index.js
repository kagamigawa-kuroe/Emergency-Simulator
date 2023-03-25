import Vuex from 'vuex'

export const urgences = new Vuex.Store({
    state(){
        return{
            //heure
            hour: 0,
            //minute
            minute: 0,
            //Semaine
            week : 1,
            //Jour
            day : 1,
            //Patients deja traites
            patientsTraites : 0,
            //Temps total de traitement dans chambre 1
            time1 : 0,
            //Progres
            progress1 : 10,
            //waiting list
            waitingList: [],
            //color
            color1 : '#ffa64d',
            color2 : '#aaaaaa',
            //Nombre de medecin
            nbMedecin : 10,
            //List de patient
            allPatList:[
                {
                    id:1,
                    name:'CUS1',
                    avatar: require('./379339-512.png'),
                    state: 'waiting',
                    patience: 100,
                    come:false,
                    timer: null
                },
                {
                    id:2,
                    name:'CUS2',
                    avatar: require('./379444-512.png'),
                    state: 'waiting',
                    patience: 100,
                    come:false,
                    timer: null
                },
                {
                    id:3,
                    name:'CUS3',
                    avatar: require('./379446-512.png'),
                    state: 'waiting',
                    patience: 100,
                    come:false,
                    timer: null
                },
                {
                    id:4,
                    name:'CUS4',
                    avatar: require('./379448-512.png'),
                    state: 'waiting',
                    patience: 100,
                    come:false,
                    timer: null
                },
                {
                    id:5,
                    name:'CUS5',
                    avatar: require('./iconfinder_Boss-3_379348.png'),
                    state: 'waiting',
                    patience: 100,
                    come:false,
                    timer: null
                },
                {
                    id:6,
                    name:'CUS6',
                    avatar: require('./iconfinder_Man-16_379485.png'),
                    state: 'waiting',
                    patience: 100,
                    come:false,
                    timer: null
                },
                {
                    id:7,
                    name:'CUS7',
                    avatar: require('./iconfinder_Rasta_379441.png'),
                    state: 'waiting',
                    patience: 100,
                    come:false,
                    timer: null
                },
                {
                    id:8,
                    name:'CUS8',
                    avatar: require('./379339-512.png'),
                    state: 'waiting',
                    patience: 100,
                    come:false,
                    timer: null
                },
                {
                    id:9,
                    name:'CUS9',
                    avatar: require('./379444-512.png'),
                    state: 'waiting',
                    patience: 100,
                    come:false,
                    timer: null
                },
                {
                    id:10,
                    name:'CUS10',
                    avatar: require('./379446-512.png'),
                    state: 'waiting',
                    patience: 100,
                    come:false,
                    timer: null
                },
                {
                    id:11,
                    name:'CUS11',
                    avatar: require('./379448-512.png'),
                    state: 'waiting',
                    patience: 100,
                    come:false,
                    timer: null
                },
                {
                    id:12,
                    name:'CUS12',
                    avatar: require('./iconfinder_Boss-3_379348.png'),
                    state: 'waiting',
                    patience: 100,
                    come:false,
                    timer: null
                },
                {
                    id:13,
                    name:'CUS13',
                    avatar: require('./iconfinder_Man-16_379485.png'),
                    state: 'waiting',
                    patience: 100,
                    come:false,
                    timer: null
                },
                {
                    id:14,
                    name:'CUS14',
                    avatar: require('./iconfinder_Rasta_379441.png'),
                    state: 'waiting',
                    patience: 100,
                    come:false,
                    timer: null
                },
                {
                    id:15,
                    name:'CUS15',
                    avatar: require('./iconfinder_Rasta_379441.png'),
                    state: 'waiting',
                    patience: 100,
                    come:false,
                    timer: null
                },
                {
                    id:16,
                    name:'CUS16',
                    avatar: require('./iconfinder_Man-16_379485.png'),
                    state: 'waiting',
                    patience: 100,
                    come:false,
                    timer: null
                },
                {
                    id:17,
                    name:'CUS17',
                    avatar: require('./iconfinder_Boss-3_379348.png'),
                    state: 'waiting',
                    patience: 100,
                    come:false,
                    timer: null
                },
                {
                    id:18,
                    name:'CUS18',
                    avatar: require('./379448-512.png'),
                    state: 'waiting',
                    patience: 100,
                    come:false,
                    timer: null
                },
                ],
            //Patient dans chaque salle
            cntPatient: [0,0,0,0],
            cntInfirmier: 3
        }
    },
    mutations: {
        initState() {
          this.commit('startDataUpdating')
          this.commit('resumeAddWaiting')
        },
        // 开启日期时间变换
        startDataUpdating(state) {
          // 每天240分钟 每天4小时
          state.timer = setInterval(() => {
            if (state.minute < 60) {
              state.minute++
            } else if (state.hour < 3) {
              state.minute = 0
              state.hour++
            } else if (state.day < 7) {
                this.commit('resetAllPat') //一天结束，重置顾客状态
              state.minute = 0
              state.hour = 0
              state.day += 1
            } else {
              state.minute = 0
              state.hour = 0
              state.day = 1
              state.week += 1
            }
          }, 500)
        },
        // 重置所有顾客的状态
    resetAllPat(state){
        state.allPatList.forEach((cus) => {
          cus.come = false
          cus.patience = 100
          cus.state = 'waiting'
          cus.timer = null
        })
      },
      // 等待队列随机时间添加人物
      resumeAddWaiting(state) {
        let len = state.allPatList.length
        state.waitAddTimer = setTimeout(function addWait() {
          let cus = state.allPatList[Math.floor(Math.random() * len)]
          if (state.waitingList.length < 5) {
            if (cus.come === false){
              let wt = JSON.parse(JSON.stringify(cus))
              state.waitingList.push(wt)
              wt.timer = setTimeout(function time() {
                wt.patience--
                if (wt.patience > 0) {
                  wt.timer = setTimeout(time, 100)
                } else {
                  let idx = state.waitingList.findIndex((item) => {
                    return item.id === wt.id
                  })
                  state.waitingList.splice(idx, 1)
                }
              }, 100)
            }
          }
          cus.come = true
          state.waitAddTimer = setTimeout(addWait, 4000)
        }, 3000)
      }
    },
    //Fonction pour le pretraitement des infirmiers
    ordering(state) {
        if (state.cntInfirmier == 0) {
            this.commit('prompting', ["Il n'existe pas d'infirmier disponible", 'negative'])
            return
        }
        this.commit('pause')
        state.orderingCus = state.waitingList[0]
        // 从等待队列中删除第一个顾客
        state.waitingList.splice(0, 1)
        state.menuing = true
    },

})
