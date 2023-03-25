<template xmlns:progress="http://java.sun.com/xml/ns/javaee">
  <div class = "app"  v-bind:items="items">
    <header style="margin-bottom: 20px">
<!--      <img class="people-icon" style="width: 100%;" src="src/assets/images/ambulance.png" alt=""/>-->
      <div class="header-item date" :style="progress">
        <span>Week{{week}} </span> &nbsp; <span>Day{{day}}</span>
      </div>
      <div class="header-item patients">
        Salle des urgences
      </div>

      <div class="image-box">
        <img class="image" src="./assets/images/ambulance.png">
      </div>
<!-- ../store/379339-512.png patient
     ../store/379339-512.png medecin
-->
        <div style="display: flex; flex-wrap: wrap">
          <div class="waiting-area" style="float: right">

            <div class="wait" style="width: 60%">
              <div class="waiter" v-for=" i in (items[5][0])" :key="i">
                <div class="people p2 waiting" >
<!--                     :class="wt.state" >-->
                  <img class="people-icon" style="width: 100%;" src="../store/379339-512.png" alt=""/>
<!--                  <div class="label" :style="calculateWaitStyle(wt)">-->
                    Waiting
<!--                  </div>-->
                </div>
              </div>
            </div>
          </div>
        <div class = "input-box" style="margin-top: 5%;margin-left: 0; margin-right: auto ">
          <div class="maladie"><label for="country">La maladie du nouveau patient</label></div>
          
          <select v-model = "maladie"  style="width: 100%">
            <option value=1>Rhume</option>
            <option value=1>Fievre</option>
            <option value=2>Mal à l'estomac</option>
            <option value=2>Diarrhée</option>
            <option value=3>Rupture des os</option>
<!--            心绞痛-->
            <option value=3>Angine</option>
<!--            中风-->
            <option value=4>Accident vasculaire cérébral</option>
            <option value=4>Cracher du sang</option>
<!--            休克-->
            <option value=5>Choc</option>
<!--            心脏骤停-->
            <option value=5>Arrêt cardiaque</option>
          </select>
          <div class="button" @click="createPatient" style="margin-bottom: auto">Ajouter</div>
        </div>
        </div>

      <div>
        <div >
          <div style="display: flex; flex-wrap: wrap">
            <div style="width:33.3%; margin-left: 0;">
              <div class="Accueil ac1" v-if="nbAccueil>=1" style="margin-top: 2%;height: 6rem; width:35%; float: left;">
                Infirmier 1
                <div class="people p1 complete" v-if="items[0][0]==1">
                  <img style="width: 100%;" src="../store/379339-512.png" class="people-icon">
                  <div class="redC"></div>
                  <div class="redR"></div>
                </div>
              </div>
                <!--              <div class="people cus" :class="cus.state">-->
  <!--                <img class="people-icon" style="width: 100%" :src="cus.avatar"/>-->
  <!--              </div>-->
  
            
              <div class="Accueil ac2" v-if="nbAccueil>=2" style="margin-top: 2%;height: 6rem; width:35%;float: left">Infirmier 2
                <div class="people p1 complete" v-if="items[0][1]==1">
                  <img style="width: 100%;" src="../store/379339-512.png" class="people-icon">
                <!--                 src="/store/iconfinder_Chef-2_379358.png" alt=""/>-->
                <div class="redC"></div>
                <div class="redR"></div>
              </div>
              </div>
              <div class="Accueil ac3" v-if="nbAccueil>=3" style="margin-top: 2%;height: 6rem; width:35%;float: left">Infirmier 3
                <div class="people p1 complete" v-if="items[0][2]==1">
                  <img style="width: 100%;" src="../store/379339-512.png" class="people-icon">
                <!--                 src="/store/iconfinder_Chef-2_379358.png" alt=""/>-->
                <div class="redC"></div>
                <div class="redR"></div>
              </div>
              </div>
              <div class="Accueil ac4" v-if="nbAccueil>=4" style="margin-top: 2%;height: 6rem; width:35%;float: left">Infirmier 4
                <div class="people p1 complete" v-if="items[0][3]==1">
                  <img style="width: 100%;" src="../store/379339-512.png" class="people-icon">
                <!--                 src="/store/iconfinder_Chef-2_379358.png" alt=""/>-->
                <div class="redC"></div>
                <div class="redR"></div>
              </div>
              </div>
              <div class="Accueil ac5" v-if="nbAccueil>=5" style="margin-top: 2%;height: 6rem; width:35%;float: left">Infirmier 5
                <div class="people p1 complete" v-if="items[0][4]==1">
                  <img style="width: 100%;" src="../store/379339-512.png" class="people-icon">
                <!--                 src="/store/iconfinder_Chef-2_379358.png" alt=""/>-->
                <div class="redC"></div>
                <div class="redR"></div>
              </div>
              </div>
              <div class="button activerAccueil" @click="activerAccueil" style="float: left; margin-top: auto; margin-bottom: auto">➕</div>
              <div class="button desactiverAccueil" @click="desactiverAccueil"  style="float: left; margin-top: auto; margin-bottom: auto">➖</div>
            </div>
            <div style="width:33.3%; margin-left: 0;justify-content:center">
              <div class="Accueil nbMedecin" style="vertical-align:middle; margin-right: auto; margin-left: auto;background: #ffa64d;height: 20rem; width:100%">
              <div>
                <img style="width: 9%;display:inline;" src="../store/379444-512.png" class="people-icon">
                <p style="display:inline;"><font size=4>Nombre de medecin level 1:{{items[4][0]}}</font></p>
              </div>
              <div>
                <img style="width: 9%;display:inline;" src="../store/379444-512.png" class="people-icon">
                <p style="display:inline;"><font size=4>Nombre de medecin level 2:{{items[4][1]}}</font></p>
              </div>
              <div>
                <img style="width: 9%;display:inline;" src="../store/379444-512.png" class="people-icon">
                <p style="display:inline;"><font size=4>Nombre de medecin level 3:{{items[4][2]}}</font></p>
              </div>
              <div>
                <img style="width: 9%;display:inline;" src="../store/379444-512.png" class="people-icon">
                <p style="display:inline;"><font size=4>Nombre de medecin level 4:{{items[4][3]}}</font></p>
              </div>
              <div>
                <img style="width: 9%;display:inline;" src="../store/379444-512.png" class="people-icon">
                <p style="display:inline;"><font size=4>Nombre de medecin level 5:{{items[4][4]}}</font></p>
              </div>
                
                
                  <div class="input box" style="">
                    <input v-model="levelMedecin" type="number" style="width:5rem; height:2rem; float:left" min=1 max=5>
                   
                    
                  </div>
                  <div class="button activerDoc" @click="activerDoc" style="float:left" >➕</div>
                    <div class="button desactiverDoc" @click="desactiverDoc" style="float:left">➖</div>
                  
                  
                
              
          </div>

              
              
            </div>
            <div style="width:33.3%; margin-right: 0;">
              <div class="Accueil infi1" v-if="nbInfirmier>=1" style="margin-top:1.5%;height: 5.5rem;float:right;display: block;width:35%">Accueil 1
                <div class="people p1 complete" v-if="items[1][0]==1">
                  <img style="width: 100%;" src="../store/379339-512.png" class="people-icon">
                  <!--                 src="/store/iconfinder_Chef-2_379358.png" alt=""/>-->
                  <div class="redC"></div>
                  <div class="redR"></div>
                </div>
                
              </div>
  
              <div class="Accueil infi2" v-if="nbInfirmier>=2" style="margin-top:1.5%;height:5.5rem;float: right; display: block;width:35%">Accueil 2
                <div class="people p1 complete" v-if="items[1][1]==1">
                  <img style="width: 100%;" src="../store/379339-512.png" class="people-icon">
                  <!--                 src="/store/iconfinder_Chef-2_379358.png" alt=""/>-->
                  <div class="redC"></div>
                  <div class="redR"></div>
                </div>
                
              </div>
  <!--            <div class="Accueil infi2" v-else style="margin-top:1.5%;height:3rem; {background: #aaaaaa}" >Reception Infirmiere 2</div>-->
  
              <div class="Accueil infi3" v-if="nbInfirmier>=3" style="margin-top:1.5%;height: 5.5rem;float: right;display: block;width:35%">Accueil 3
                <div class="people p1 complete" v-if="items[1][2]==1">
                  <img style="width: 100%;" src="../store/379339-512.png" class="people-icon">
                  <!--                 src="/store/iconfinder_Chef-2_379358.png" alt=""/>-->
                  <div class="redC"></div>
                  <div class="redR"></div>
                </div>
                
              </div>
  
              <div class="Accueil infi4" v-if="nbInfirmier>=4" style="margin-top:1.5%;height: 5.5rem;float:right;display: block;width:35%">Accueil 4
                <div class="people p1 complete" v-if="items[1][3]==1">
                  <img style="width: 100%;" src="../store/379339-512.png" class="people-icon">
                  <!--                 src="/store/iconfinder_Chef-2_379358.png" alt=""/>-->
                  <div class="redC"></div>
                  <div class="redR"></div>
                </div>
                
              </div>
  
              <div class="Accueil infi5" v-if="nbInfirmier>=5" style="margin-top: 1.5%;height: 5.5rem;float:right;display: block;width:35%">Accueil 5
                <div class="people p1 complete" v-if="items[1][4]==1">
                  <img style="width: 100%;" src="../store/379339-512.png" class="people-icon">
                  <!--                 src="/store/iconfinder_Chef-2_379358.png" alt=""/>-->
                  <div class="redC"></div>
                  <div class="redR"></div>
                </div>
                
              </div>
              <div class="button activerInfirmier" @click="activerInfirmier" style="float: right">➕</div>
              <div class="button desactiverInfirmier" @click="desactiverInfirmier" style="float: right">➖</div>
            </div>
          </div>
        </div>

      </div>

    </header>

    <main style="margin-top: 20px; width: 100%;">

        <div class="salleAttente" style="width: 100%; margin-top: 2rem;display: flex; flex-wrap: wrap">
          Salle d'attente
          
          <div class="people p2 complete" v-for="i in items[2][0]" :key="i" >
            <img style="width: 100%;" src="../store/379339-512.png" class="people-icon">

            <!--          <div v-for="i in nbPatient">-->
<!--            <div class="people cus" :class="cus.state">-->
<!--              <img class="people-icon" style="width: 100%" :src="cus.avatar"/>-->
<!--            </div>-->

<!--          </div>-->
        </div>
      </div>

      <div style="display: flex; flex-wrap: wrap">
        <div style="width:100%; margin-left: 0;">
          <div class="Room Room1" v-if="nbSalle>=1" style="float:left;">
            Room1-level{{items[6][0]}}
            <div class="people p1 complete" v-if="items[3][0]==1">
              <img style="width: 100%;" src="../store/379339-512.png" class="people-icon">
              <!--                 src="/store/iconfinder_Chef-2_379358.png" alt=""/>-->
              <div class="redC"></div>
              <div class="redR"></div>
            </div>
          </div>
          <div class="Room Room2" v-if="nbSalle>=2" style="float:left;margin-left: 2rem">
            Room2-level{{items[6][1]}}
            <div class="people p1 complete" v-if="items[3][1]==1">
              <img style="width: 100%;" src="../store/379339-512.png" class="people-icon">
              <!--                 src="/store/iconfinder_Chef-2_379358.png" alt=""/>-->
              <div class="redC"></div>
              <div class="redR"></div>
            </div>
          </div>
          <div class="Room Room3" v-if="nbSalle>=3" style="float:left;margin-left: 2rem">
            Room3-level{{items[6][2]}}
            <div class="people p1 complete" v-if="items[3][2]==1">
              <img style="width: 100%;" src="../store/379339-512.png" class="people-icon">
              <!--                 src="/store/iconfinder_Chef-2_379358.png" alt=""/>-->
              <div class="redC"></div>
              <div class="redR"></div>
            </div>
          </div>
          <div class="Room Room4" v-if="nbSalle>=4" style="float:left;margin-left: 2rem">
            Room4-level{{items[6][3]}}
            <div class="people p1 complete" v-if="items[3][3]==1">
              <img style="width: 100%;" src="../store/379339-512.png" class="people-icon">
              <!--                 src="/store/iconfinder_Chef-2_379358.png" alt=""/>-->
              <div class="redC"></div>
              <div class="redR"></div>
            </div>
          </div>
          <div class="Room Room5" v-if="nbSalle>=5" style="float:left;margin-left: 2rem">
            Room5-level{{items[6][4]}}
            <div class="people p1 complete" v-if="items[3][4]==1">
              <img style="width: 100%;" src="../store/379339-512.png" class="people-icon">
              <!--                 src="/store/iconfinder_Chef-2_379358.png" alt=""/>-->
              <div class="redC"></div>
              <div class="redR"></div>
            </div>
          </div>
          <div class="Room Room6" v-if="nbSalle>=6" style="float:left;margin-left: 2rem">
            Room6-level{{items[6][5]}}
            <div class="people p1 complete" v-if="items[3][5]==1">
              <img style="width: 100%;" src="../store/379339-512.png" class="people-icon">
              <!--                 src="/store/iconfinder_Chef-2_379358.png" alt=""/>-->
              <div class="redC"></div>
              <div class="redR"></div>
            </div>
          </div>
          <div class="Room Room7" v-if="nbSalle>=7" style="float:left;margin-left: 2rem">
            Room7-level{{items[6][6]}}
            <div class="people p1 complete" v-if="items[3][6]==1">
              <img style="width: 100%;" src="../store/379339-512.png" class="people-icon">
              <!--                 src="/store/iconfinder_Chef-2_379358.png" alt=""/>-->
              <div class="redC"></div>
              <div class="redR"></div>
            </div>
          </div>
          <div class="Room Room8" v-if="nbSalle>=8" style="float:left;margin-left: 2rem">
            Room8-level{{items[6][7]}}
            <div class="people p1 complete" v-if="items[3][7]==1">
              <img style="width: 100%;" src="../store/379339-512.png" class="people-icon">
              <!--                 src="/store/iconfinder_Chef-2_379358.png" alt=""/>-->
              <div class="redC"></div>
              <div class="redR"></div>
            </div>
          </div>
          <div class="Room Room9" v-if="nbSalle>=9" style="float:left;margin-left: 2rem">
            Room9-level{{items[6][8]}}
            <div class="people p1 complete" v-if="items[3][8]==1">
              <img style="width: 100%;" src="../store/379339-512.png" class="people-icon">
              <!--                 src="/store/iconfinder_Chef-2_379358.png" alt=""/>-->
              <div class="redC"></div>
              <div class="redR"></div>
            </div>
          </div>
          <div class="Room Room10" v-if="nbSalle>=10" style="float:left;margin-left: 2rem">
            Room10-level{{items[6][9]}}
            <div class="people p1 complete" v-if="items[3][9]==1">
              <img style="width: 100%;" src="../store/379339-512.png" class="people-icon">
              <!--                 src="/store/iconfinder_Chef-2_379358.png" alt=""/>-->
              <div class="redC"></div>
              <div class="redR"></div>
            </div>
          </div>
        </div>
      </div>
      <!-- 按钮 -->
      <div style="height: 2rem">
        <div class="input box">
          <input v-model="levelSalle" type="number" style="width:5rem; height:2rem; float: left" min=1 max=5>
        </div>
        <div class="button activerSalle" @click="activerSalle" style="float: left">➕</div>
        <div class="button desactiverSalle" @click="desactiverSalle" style="float: left">➖</div>
      </div>

        
          
          
      


    </main>
  </div>
  <welcome></welcome>
  <img ref="conf0" :src="backgroundImage" style="position:absolute;width: 100%;height: 100%;left: 0;top: 0;z-index: -1" alt="" >

</template>

<script>
import Welcome from "./components/welcomePop";
import {mapState} from 'vuex'
import axios from 'axios'

export default {
  name: 'App',
  components: {
    Welcome
  },
  data(){
    return{
      //Les infos ensemble passees
      items:
          [
            //Accueil
            [],
            //Infirmier
            [],
            //Nb patients
            [],
            //Room
            [],
            //Nb medecin
            [],
            //File d'attente
            [],
            // level de room
            []


            //Accueil
          // [0,1,1],
          // //Infirmier
          // [1,0,0],
          // //Nb patients
          // [3],
          // //Room
          // [1,1,0,0,1],
          // //Nb medecin
          // [1]
      ],
      backgroundImage: require('/src/assets/images/2382727.jpg'),
      backColorMap: {
        free: 'background: linear-gradient(to right, #dddddd 50%, #aaaaaa 50%);',
        appease: 'background: linear-gradient(to right, #661a00 50%, #401000 50%);',
        safe: 'background:  #00b200;',
        warning: 'background: linear-gradient(to right,#ff9122 50%,#d96d00  50%);',
        danger: 'background: linear-gradient(to right,#ff2020 50%, #b20000 50%);',
        destroy: 'background-color: #535362;text-decoration: line-through;'
      },
      maladie : 0,
      nbAccueil : 3,
      nbInfirmier : 3,
      nbSalle : 5,
      nbMedecin:9,
      levelMedecin:0,
      levelSalle:0,
      nbPatient:0,
      Interval : null
    }
  },
  
  // Fonction pour rafraichir les infos toutes les 0.5 secondes
  mounted() {
    // let interval= null
    // const instance = axios({
    //       baseURL:'http://localhost:8082/getinfo',
    //       timeout :5000
    //     })

    this.Interval=setInterval(() => {
    //   instance
    //       .get('http://localhost:8082/getinfo')
    //       .then(response => {
    //         this.$nextTick(() => {
    //           // 修改组件数据
    //           this.items = response.data;
    //         });
    //         console.log(response.data);})
    //       .catch(e => {
    //         console.log(e);
    //       });
    // }, 500);
      axios
          .get('http://localhost:8082/getinfo')
          .then(response => {
            this.$nextTick(() => {
              // 修改组件数据
              this.items = response.data;
            });
            console.log(response.data);})
          .catch(e => {
                  console.log(e);
          });
    }, 100);
    // this.beforeUnmount()
  },
  beforeUnmount() {
    clearInterval(this.Interval);
  },

  methods: {
    //Fonction d'essai pour pousser l'info du patient au backend
  
    createPatient() {
      axios
          .get('http://127.0.0.1:8082/createPatient?test='+this.maladie.toString())
          .then(response => {
            // 创建成功，将新用户添加到用户列表中
            console.log(response.data);
      })
          .catch(e => {
        console.log(e);
      });
    },

    // 计算顾客耐心值
    calculateWaitStyle(wt) {
      return 'background: linear-gradient(to right, #006dd9 ' + (100 - wt.patience) + '%, #2693ff ' + (100 - wt.patience) + '%);'
    },
    activerAccueil(){
      if (this.nbAccueil == 5){
        this.nbAccueil = 5
      }else
        this.nbAccueil ++
      axios.post('http://localhost:8082/activerAccueil', null)
          .then(response => {
            console.log(response.data)
          });

    },

    desactiverAccueil(){
      if (this.nbAccueil == 1){
        this.nbAccueil = 1
      }else
        this.nbAccueil --
      axios.post('http://localhost:8082/desactiverAccueil', null)
          .then(response => {
            console.log(response.data)
          });

    },
    activerInfirmier(){
      if (this.nbInfirmier == 5){
        this.nbInfirmier = 5
      }else
        this.nbInfirmier ++
      axios.post('http://localhost:8082/activerInfirmier', null)
          .then(response => {
            console.log(response.data)
          });

    },

    desactiverInfirmier(){
      if (this.nbInfirmier == 1){
        this.nbInfirmier = 1
      }else
        this.nbInfirmier --
      axios.post('http://localhost:8082/desactiverInfirmier', null)
          .then(response => {
            console.log(response.data)
          });

    },

    activerSalle(){
      // alert(this.levelSalle.toString())
      if (this.nbSalle == 10){
        this.nbSalle = 10
      } else
        this.nbSalle ++
      axios.post('http://localhost:8082/activerSalle?test='+this.levelSalle.toString())
      // axios.post('http://localhost:8082/activerSalle'+ this.levelSalle.toString())
          .then(response => {
            console.log(response.data)
          });

    },

    desactiverSalle(){
      if (this.nbSalle == 1){
        this.nbSalle = 1
      }else
        this.nbSalle --
      axios.post('http://localhost:8082/desactiverSalle', null)
          .then(response => {
            console.log(response.data)
          });

    },
    activerDoc(){
        this.nbMedecin ++
      axios.get('http://127.0.0.1:8082/activerDoc?test='+this.levelMedecin.toString())
          .then(response => {
            // 创建成功，将新用户添加到用户列表中
            console.log(response.data);
          })
          .catch(e => {
            console.log(e);
          });


    },
    desactiverDoc(){
        if (this.nbMedecin == 1){
        this.nbMedecin = 1
      }else
        this.nbMedecin --
      axios.post('http://localhost:8082/desactiverDoc?test='+this.levelMedecin.toString())
          .then(response => {
            console.log(response.data)
          });

    },
  },

  computed: {
    ...mapState([
        'hour',
        'minute',
        'day',
        'week',
        'patientsTraites',
        'time1',
        'progress1',
        'waitingList',
        'allPatListList',
        'cntInfirmier',
        'color1',
        'color2',
        'nbMedecin'
    ]),
    progress(){
      let pro = (this.hour*60+this.minute)/240*100
      return 'background: linear-gradient(to right, #ffd24d '+pro+'%, #ffe699 '+pro+'%);'
    },

    // name:'createPatient',
    // data(){
    //   return{
    //     gravite : 0
    //   }
    // },
  }
}
</script>

<style>
#app {
  margin: 0;
  padding:0;
  width: 100%;
  height: 100%;
  box-sizing: border-box;
}

.header-item {
  background: linear-gradient(to bottom, #ffe699 50%, #ffd24d 50%);
  line-height: 1.8rem;
  width: 20%;
  border: 0.3rem solid #8c6900;
  border-radius: 99999px;
  font-size: 1rem;
  /*display: inline-flex;*/
  position: relative;
  margin-bottom: 10%;
}
/*Le temps passe*/
.date{
  margin-left:auto;
  margin-right:auto;
  position: absolute;
  right: 0;
}

/*Le nombre de patients traites*/
.patients{
  margin-left:auto;
  margin-right:auto;
  position: absolute;
  left: 0;
}


.image-box{
  height: 100px;
  width: 100px;
  background-color: transparent;
  margin-left: auto;
  margin-right: auto;
}
.image {
  height: 100px;
  width: 100px;
  /*background-image: url("assets/images/ambulance.png");*/
  background-repeat: no-repeat;
  max-width: 100%;
  max-height: 100%;
  background-size: contain;
  display: inline-block;
  /*background-position: top;*/
}

/*Room*/
main .Room {
  background-color: #ffa64d;
  margin: auto;
  height: 5rem;
  width: 8%;  /*left: 50%;*/
  /*top: 70%;*/
  /*float: left;*/
  border-radius: 0.8rem;
  border: 0.3rem solid white;
  box-sizing: border-box;
  display: inline-block;
  /*flex-wrap: wrap;*/
  justify-content: left;
  align-items: flex-start;
  text-align: center;
  margin-bottom: 1%;
  /*padding-bottom: 2rem;*/
  /*display: inline-flex;*/

}

/*main .Room1 {*/
/*  float: left;*/
/*}*/

/*main .Room3 {*/
/*  float: left;*/
/*}*/

/*main .Room2 {*/
/*  float: right;*/
/*}*/

/*main .Room4 {*/
/*  float: right;*/
/*}*/
.buttonArea{
  margin-top: 3rem;
  width: 40%;
  display: inline-block;
}
 .Accueil {
  /*margin-left: auto;*/
  /*margin-left: auto;*/
  /*align-items: center;*/
  /*margin-top: 50px;*/
  margin-left: auto;
  margin-right: auto;
  background-color: #ffd24d;
  height: 10rem;
  width: 40%;
  /*left: 50%;*/
  /*top: 70%;*/
  /*float: left;*/
  border-radius: 0.8rem;
  border: 0.3rem solid white;
  box-sizing: border-box;
  display: flex;
  flex-wrap: wrap;
  justify-content: left;
  align-items: normal;
  text-align: start;
  padding-bottom: 1.5rem;

}
 .salleAttente{
   background: chocolate;
   margin: auto;
   height: 5rem;
   width: 10%;  /*left: 50%;*/
   /*top: 70%;*/
   /*float: left;*/
   border-radius: 0.8rem;
   border: 0.3rem solid white;
   box-sizing: border-box;
   display: inline-block;
   /*flex-wrap: wrap;*/
   justify-content: left;
   align-items: flex-start;
   text-align: start;
   margin-bottom: 1%;

 }

 .activerAccueil{
   margin-top: 3rem;
 }

 .desactiverAccueil{
   margin-left: 1rem;
 }

 .activerDoc{
   height: 12%;
 }
 .desactiverDoc{
  height: 12%;
}

 /*Espace pour les entrees*/
.input-box{
  height: 80px;
  width: 20%;
  background-color: #ffd24d;
  margin-left: auto;
  margin-right: auto;
  border-radius: 0.8rem;
  border: 0.3rem solid white;
  padding-bottom: 1.5rem;
}

/*Espaces pour renseigner les infos du noveau patient*/
input[type=text], select {
  width: 40%;
  padding: 12px 20px;
  margin: 8px 0;
  display: inline-block;
  border: 1px solid #ccc;
  border-radius: 4px;
  box-sizing: border-box;
  align-items: center;
  margin-left: auto;
  margin-right: auto;

}
/*Bouton pour envoyer les infos*/
input[type=submit] {
  width: 20%;
  background-color: #4CAF50;
  color: white;
  padding: 14px 20px;
  margin: 8px 0;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  float: right;
}

/*barre de progression*/
.progress-bar{
  height: 2rem;
  background-color: #ac91ff;
}
/*Zone d'attente*/
.waiting-area {
  position: relative;
  width: 80%;
  margin: 0 auto;
}

.waiting-area .wait {
  display: flex;
  flex-direction: row-reverse;
  position: absolute;
  right: 0;
  transform: translateY(-50%);
}

.waiting-area .wait .waiter {
  width: 3.5rem;
  transform: translateX(-50%);
}


/*.button{*/
/*  display: flex;*/
/*  justify-content: center;*/
/*  align-items: center;*/
/*  !*float: center;*!*/
/*  width: 30%;*/
/*  margin-left:auto;*/
/*  margin-right:auto;*/
/*}*/


</style>
