import { BehaviorSubject } from "rxjs";
import { default as ws, WebSocketServer } from "ws";
import { RaceLap } from "../models/race.js";

const socketPool = new Set<ws>();

export async function init(bus: BehaviorSubject<RaceLap>) {
  const server = new WebSocketServer({ port: 3000 });

  server.on("connection", (socket) => {
    socketPool.add(socket);

    socket.on("close", () => {
      socketPool.delete(socket);
    });
  });

  bus.subscribe((raceLap: RaceLap) => {
    const payload = JSON.stringify(raceLap);

    socketPool.forEach((socket: ws) => {
      socket.send(payload);
    });
  });
}
