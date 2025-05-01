import Room from "@/api/rooms";
import Auth from "@/api/auth";
import User from "@/api/users"

class API {
  public api: object = {
    rooms: new Room(),
    auth: new Auth(),
    users: new User(),
  }
}

export default API;
