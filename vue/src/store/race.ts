import { defineStore } from "pinia";
import { Ref, ref } from "vue";
import { RaceLap, RaceResult } from "../models/race";

export const useRaceStore = defineStore("race", () => {
  const result: Ref<RaceResult | null> = ref(null);
  const laps: Ref<RaceLap[]> = ref([]);

  const socket = new WebSocket("ws://localhost:8082");
  socket.addEventListener("message", (event: MessageEvent<string>) => {
    try {
      addLap(JSON.parse(event.data) as RaceLap);
    } catch (e) {
      console.error("Lap error", e);
    }
  });

  function setResult(newResult: RaceResult) {
    result.value = newResult;
    laps.value = [];
  }
  function addLap(lap: RaceLap) {
    if (
      !result.value ||
      result.value.season !== lap.season ||
      result.value.name !== lap.name
    ) {
      throw "Lap race different from result one";
    }
    laps.value = [...laps.value, lap];
  }
  return { result, laps, setResult, addLap };
});
