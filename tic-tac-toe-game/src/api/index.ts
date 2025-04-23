import Rooms from "@/api/rooms";

class API {
  public api: object = {
    rooms: new Rooms(),
  }
}

export default new API().api;
