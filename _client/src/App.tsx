import { useCallback } from "react";
import { FormProvider, useForm } from "react-hook-form";
import { InputText } from "@/src/layers/shared/components/Input/Text/InputText.tsx";
import {
  User_Create_Input,
  User_Login_Input,
} from "@/src/layers/features/users/types.ts";
import { UserCreateInputValidation } from "@/src/layers/features/users/validation/UserCreateInputValidation.ts";
import { useDebug } from "@/src/layers/features/debug/hooks/useDebug.tsx";
import { zodResolver } from "@hookform/resolvers/zod";
import { useUserCreate } from "@/src/layers/features/users/hooks/useUserCreate.tsx";
import { useUser } from "@/src/layers/features/users/hooks/useUser.tsx";
import "@/src/layers/network/endpoints.ts";
import { SnackbarProvider } from "notistack";
import { useUserLogin } from "@/src/layers/features/users/hooks/useUserLogin.tsx";
import { UserLoginInputValidation } from "@/src/layers/features/users/validation/UserLoginInputValidation.ts";
import { useUserLogout } from "@/src/layers/features/users/hooks/useUserLogout.tsx";

function App() {
  // props
  const debug = useDebug();
  const { localUser, localUserStatus, localUserNetworkError, localUserToken } =
    useUser();
  const { userCreate } = useUserCreate();
  const { userLogin } = useUserLogin();
  const { userLogout } = useUserLogout();

  // rhf
  const userCreateFormMethods = useForm<User_Create_Input>({
    resolver: zodResolver(UserCreateInputValidation),
  });
  const userLoginFormMethods = useForm<User_Login_Input>({
    resolver: zodResolver(UserLoginInputValidation),
  });

  // functions
  const onFormSubmit_Create = useCallback(
    (data: User_Create_Input) => {
      debug.log("onFormSubmit", data);
      userCreate(data).then();
    },
    [debug, userCreate],
  );

  const onFormSubmit_Login = useCallback(
    (data: User_Login_Input) => {
      debug.log("onFormSubmit", data);
      userLogin(data).then();
    },
    [debug, userLogin],
  );

  return (
    <>
      <SnackbarProvider />

      <h1>Hello, World!</h1>

      <div>
        <pre>{JSON.stringify(localUser, null, 4)}</pre>
        <pre>{JSON.stringify(localUserStatus, null, 4)}</pre>
      </div>

      <div>
        <FormProvider {...userCreateFormMethods}>
          <form
            onSubmit={userCreateFormMethods.handleSubmit(onFormSubmit_Create)}
          >
            <div>
              <label>Username:</label>
              <InputText
                name={"username"}
                placeholder={"enter username here"}
              />
            </div>

            <div>
              <label>Email:</label>
              <InputText name={"email"} placeholder={"enter email here"} />
            </div>

            <div>
              <label>Password:</label>
              <InputText
                name={"password"}
                isPassword={true}
                placeholder={"enter password here"}
              />
            </div>

            <div>
              <button type={"submit"}>
                {localUserStatus === "idle" ? "Create account" : "Loading..."}
              </button>
            </div>
          </form>
        </FormProvider>
      </div>

      <div>
        <FormProvider {...userLoginFormMethods}>
          <form
            onSubmit={userLoginFormMethods.handleSubmit(onFormSubmit_Login)}
          >
            <div>
              <label>Username:</label>
              <InputText
                name={"username"}
                placeholder={"enter username here"}
              />
            </div>

            <div>
              <label>Password:</label>
              <InputText
                name={"password"}
                isPassword={true}
                placeholder={"enter password here"}
              />
            </div>

            <div>
              <button type={"submit"}>
                {localUserStatus === "idle" ? "Sign in" : "Loading..."}
              </button>
            </div>
          </form>
        </FormProvider>
      </div>

      {localUserNetworkError && (
        <div style={{ color: "red" }}>{localUserNetworkError}</div>
      )}

      {localUserToken && (
        <div>
          <p>{localUserToken}</p>
          <button
            type={"button"}
            onClick={() => {
              userLogout();
            }}
          >
            Logout
          </button>
        </div>
      )}
    </>
  );
}

export default App;
