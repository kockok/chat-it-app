/* tslint:disable */
/* eslint-disable */
// @generated
// This file was automatically generated and should not be edited.

// ====================================================
// GraphQL mutation operation: registerUser
// ====================================================

export interface registerUser_newUser_user {
  __typename: "User";
  username: string;
  id: string;
}

export interface registerUser_newUser {
  __typename: "UserSession";
  user: registerUser_newUser_user;
  session: string;
}

export interface registerUser {
  newUser: registerUser_newUser;
}

export interface registerUserVariables {
  username?: string | null;
  password?: string | null;
}
