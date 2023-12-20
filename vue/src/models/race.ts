export type RaceSelection = {
  season: string;
  raceNumber: string;
};

type Race = {
  season: string;
  name: string;
  time: Date;
};

export type RaceResult = Race & {
  nbLaps: number;
};

export type RaceLap = Race & {
  lapNumber: number;
  timings: Timing[];
};

export type Timing = {
  driverId: string;
  position: number;
  time: number;
};
