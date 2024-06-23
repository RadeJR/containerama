export class User {
  given_name: string;
  family_name: string;
  username: string;
  email: string;
  picture: string;

  constructor() {
    this.picture = "";
    this.email = "";
    this.username = "";
    this.given_name = "";
    this.family_name = "";
  }
}
