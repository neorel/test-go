import cors from "cors";
import express from "express";
import { BehaviorSubject } from "rxjs";
import * as redis from "./services/redis.js";
import * as websocket from "./services/websocket.js";
import { RaceLap } from "./models/race.js";

const bus = new BehaviorSubject<RaceLap>(null);
redis.init(bus);
websocket.init(bus);

const app = express();
app.use(cors());
app.get("/race", async (req, res) => {
  try {
    const response = await fetch(
      "http://go:8080/race?season=" +
        req.query.season +
        "&number=" +
        req.query.number
    );
    const data = await response.json();
    res.send(data);
  } catch (e) {
    res.status(500).send(e);
  }
});
app.listen(8080);
