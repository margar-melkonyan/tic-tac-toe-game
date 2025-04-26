import Room from "@/api/rooms";
import Auth from "@/api/auth";

class API {
  public api: object = {
    rooms: new Room(),
    auth: new Auth(),
  }
}

export default API;
