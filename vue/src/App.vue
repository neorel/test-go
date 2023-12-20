<template>
  <div class="app">
    <header class="app_header">
      <h1>Simulateur de course de F1</h1>
    </header>
    <aside class="app_sidebar">
      <RaceSelector @race-selected="onRaceSelected" />
    </aside>
    <main class="app_main">
      <Race :result="result" :laps="laps"></Race>
    </main>
    <footer class="app_footer">
      Créé par Aurélien LEMAITRE
    </footer>
  </div>
</template>

<script setup lang="ts">
import RaceSelector from './components/RaceSelector.vue'
import Race from './components/Race.vue';
import { RaceSelection, RaceResult } from './models/race';
import { useRaceStore } from './store/race';
import { toRefs } from 'vue';

const raceStore = useRaceStore();

const {result, laps} = toRefs(raceStore);

async function onRaceSelected(race: RaceSelection) {
  try {
    const response = await fetch("http://localhost:8081/race?season="+race.season+"&number="+race.raceNumber);
    const raceResult = (await response.json()) as RaceResult;

    raceStore.setResult(raceResult);
  } catch (e) {
    console.error(e);
  }
}
</script>

<style lang="scss">
.app {
  display: grid;
  min-height: 100%;
  width: 100%;
  grid-template-areas: "header" "sidebar" "main" "footer";
  grid-template-columns: auto auto 1fr auto;
  grid-template-rows: 1fr;
  background-color: var(--color-background);
  color: var(--color-light);

  @media (min-width: 768px) {
    grid-template-areas: "header header" "sidebar main" "footer footer";
    grid-template-columns: auto 1fr auto;
    grid-template-rows: auto 1fr;
  }
  &_header {
    grid-area: header;
    text-align: center;
    padding: 1rem;
    background: var(--color-primary);
  }
  &_sidebar {
    grid-area: sidebar;
  }
  &_main {
    grid-area: main;
  }
  &_footer {
    grid-area: footer;
    background: var(--color-primary);
    text-align: center;
    padding: 1rem;
  }
}
</style>
