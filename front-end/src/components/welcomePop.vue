<template>
    <div>
        <div v-if="welcome" id="start" class="guide">
            <div class="pop-item">
                <h3 style="font-weight: bold;text-align: center;margin-bottom: 0.3rem">Bienvenue à la simulation de la salle des urgences!</h3>
                <h3 style="font-weight: bold;text-align: center;margin-bottom: 0.3rem">Voici les fonctionnements principaux:</h3>
                <p class="task">
                  <span>Prétraitement des patients (infirmiers)</span>
                  <span>Salle d'attente</span>
                  <span>Traitement des patients (docteurs)</span>
                </p>
                <div class="button" @click="start">Démarrer</div>

            </div>

        </div>
    </div>
</template>

<script>
import {mapMutations} from 'vuex'
export default {
    data() {
        return {
            welcome: true
        }
    },
    methods:{
        ...mapMutations([
            'initState'
        ]),
        start(){
            this.welcome = false
            this.initState()
        },
    }
}
</script>

<style scoped>

    .task {
        display: flex;
        justify-content: space-between;
        align-content: center;
        line-height: 2.5rem;
    }

    .button{
      display: flex;
      justify-content: center;
      align-items: center;
      /*float: center;*/
      width: 30%;
      margin-left:auto;
      margin-right:auto;
    }

</style>