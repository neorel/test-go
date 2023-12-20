<template>
    <form class="race-selector" @submit.prevent="onSubmit">
        <h2 class="race-selector_title">Choisissez votre course</h2>
        <fieldset class="race-selector_season">
            <legend>Saison</legend>
            <select v-model="season">
                <option 
                    v-for="season in seasons" 
                    :key="season" 
                    :value="season"
                >
                    {{ season }}
                </option>
            </select>
        </fieldset>
        <fieldset class="race-selector_race-number">
            <legend>Numéro de course</legend>
            <input type="number" min="1" max="23" v-model="raceNumber"/>
        </fieldset>
        <button class="race-selector_button" type="submit">Valider</button>
    </form>
</template>

<script setup lang="ts">
import { ref } from "vue";
import { RaceSelection } from '../models/race.ts'

const emits = defineEmits<{
    raceSelected: [race: RaceSelection]
}>()

const seasons = Array.from({length: 24}, (_, index) => index + 2000)

let season = ref(2023);
let raceNumber = ref(1);

function onSubmit() {
    emits('raceSelected', {
        season: ""+season.value,
        raceNumber: ""+raceNumber.value
    })
}
</script>

<style lang="scss">
.race-selector{
    display: grid;
    flex-wrap: wrap;
    gap: 1rem;
    padding: 1rem;
    background: var(--color-dark);
    height: 100%;
    grid-template-areas: "title title" "season race-number" "button button";
    grid-template-columns: 1fr 1fr;

    @media (min-width: 768px) {
        grid-template-areas: "title" "season" "race-number" "button";
        grid-template-columns: 1fr;
        grid-template-rows: repeat(4, min-content);
    }

    &_title {
        grid-area: title;
        text-align: center;
    }
    &_season {
        grid-area: season;
    }
    &_race-number {
        grid-area: race-number;
    }
    &_button {
        all: unset;
        box-sizing: border-box;
        text-align: center;
        background-color: var(--color-primary);
        border-radius: .5rem;
        cursor: pointer;
        grid-area: button;
    }

    fieldset {
        padding: .5rem;
        border-radius: .5rem;
    }

    select, input, button {
        width: 100%;
        padding: .25rem .5rem;
    }

}
</style>