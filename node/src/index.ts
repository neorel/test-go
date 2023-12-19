import { BehaviorSubject } from "rxjs";
import * as redis from "./services/redis.js";
import * as websocket from "./services/websocket.js";
import { RaceLap } from "./models/race.js";

const bus = new BehaviorSubject<RaceLap>(null);
redis.init(bus);
websocket.init(bus);
