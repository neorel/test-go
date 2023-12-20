<template>
    <div class="race" v-if="result">
        <div class="race_season">
            Saison {{ result.season }}
        </div>
        <div class="race_name">
            {{ result.name }}
        </div>
        <div class="race_nb-laps">
            {{ result.nbLaps }} tours
        </div>
        <div class="race_detail" v-if="laps">
            <table class="race_last-lap" v-if="lastLap">
                <tr>
                    <td colspan="3">Tour {{ lastLap.lapNumber }}</td>
                </tr>
                <tr>
                    <th>Position</th>
                    <th>Pilote</th>
                    <th>Temps</th>
                </tr>
                <tr v-for="(timing, index) in lastLap.timings" :key="timing.driverId">
                    <td class="race_last-lap_position">{{ index + 1 }}</td>
                    <td class="race_last-lap_driver">{{ timing.driverId }}</td>
                    <td class="race_last-lap_time">{{ displayTime(timing.time) }}</td>
                </tr>
            </table>
        </div>
    </div>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import { RaceResult, RaceLap } from '../models/race';

const props = defineProps<{
    result?: RaceResult,
    laps: RaceLap[]
}>()

const lastLap = computed(() => props.laps.at(-1))

function displayTime(time:number) {
    let output = '';
    if (time > 60) output += Math.floor(time / 60) + ':';
    const remainder = (''+(Math.floor(time % 60 * 1000) / 1000)).split('.');
    output += remainder[0].padStart(2, '0') + '.' + remainder[1].padEnd(3, '0');
    return output;
}
</script>

<style lang="scss">
.race {
    padding: 1rem;
    max-width: 1024px;
    margin: 0 auto;
    
    &_season, &_name, &_nb-laps {
        font-size: 2rem;
    }

    &_last-lap {
        margin-top: 2rem;
        width: 100%;

        &_position {
            text-align: end;
        }
        &_driver {
            text-align: center;
        }
        &_time {
            text-align: end;
        }
    }
}
</style>

