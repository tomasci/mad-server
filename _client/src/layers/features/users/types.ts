type User = {
  id?: string;
  username?: string;
  password?: never;
  email?: string;
  created_at?: string;
  updated_at?: string;
  deleted_at?: never;
};

type User_Create_Input = {
  username: string;
  password: string;
  email: string;
};

type User_Login_Input = {
  username: string;
  password: string;
};

export type { User, User_Create_Input, User_Login_Input };
