import * as z from "zod";
import { ValidationCodes } from "@/src/layers/shared/validation/ValidationCodes.ts";

const UserCreateInputValidation = z.object({
  username: z.string().min(1, ValidationCodes.required),
  email: z.string().email(ValidationCodes.invalidEmail),
  password: z.string().min(4, ValidationCodes.stringMin),
});

export { UserCreateInputValidation };
