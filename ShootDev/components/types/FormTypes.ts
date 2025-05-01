const formTypes = ["init", "logIn", "signUp", "forgotPass", "quit"] as const;

export type FormType = (typeof formTypes)[number];
