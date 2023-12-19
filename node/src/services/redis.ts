import { createClient } from "redis";
import { BehaviorSubject } from "rxjs";
import { RaceLap } from "../models/race.js";

export async function init(bus: BehaviorSubject<RaceLap>) {
  const redisClient = createClient({
    url: "redis://redis",
  });
  redisClient.on("error", (err) => console.log(err));
  await redisClient.connect();

  redisClient.subscribe("raceLap", (message: string, channel: string) => {
    bus.next(JSON.parse(message) as RaceLap);
  });
}
