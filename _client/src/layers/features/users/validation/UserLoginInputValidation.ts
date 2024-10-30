import * as z from "zod";
import { ValidationCodes } from "@/src/layers/shared/validation/ValidationCodes.ts";

const UserLoginInputValidation = z.object({
  username: z.string().min(1, ValidationCodes.required),
  password: z.string().min(4, ValidationCodes.stringMin),
});

export { UserLoginInputValidation };
