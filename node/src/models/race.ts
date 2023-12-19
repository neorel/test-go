export type RaceLap = {
  season: string;
  name: string;
  time: Date;
  nblaps: number;
  timinigs: Timing[];
};

export type Timing = {
  driverId: string;
  position: number;
  time: number;
};
