// user base
type User = Partial<{
  id: string;
  username: string;
  password: never;
  email: string;
  created_at: string;
  updated_at: string;
  deleted_at: never;
}>;

// user token
type User_Token = string;

// create user
type User_Create_Input = Pick<User, "username" | "password" | "email">;
type User_Create_Response = User;

// login user
type User_Login_Input = Pick<User, "username" | "password">;
type User_Login_Response = {
  token: User_Token;
  user: User;
};

export type {
  User,
  User_Create_Input,
  User_Create_Response,
  User_Login_Input,
  User_Login_Response,
  User_Token,
};
