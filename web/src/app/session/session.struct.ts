export class Session {
  userId: number;
  isValid: boolean;
  username: string;

  constructor(userId: number, isValid: boolean, username: string) {
    this.userId = userId;
    this.isValid = isValid;
    this.username = username;
  }
}
