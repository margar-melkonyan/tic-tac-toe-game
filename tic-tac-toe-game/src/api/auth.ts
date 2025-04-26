import BaseAPI from "@/api/base";

class Auth extends BaseAPI {
  protected URI: string = 'api/auth';
  public urls: object = {
    signIn: (): string => `${this.baseURL}/${this.URI}/sign-in`,
    signUp: (): string => `${this.baseURL}/${this.URI}/sign-up`,
    signOut: (): string => `${this.baseURL}/${this.URI}/sign-out`,
    refreshToken: (): string => `${this.baseURL}/${this.URI}/refresh-token`
  };
}
export default Auth;
