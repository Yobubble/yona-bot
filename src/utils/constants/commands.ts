import { CommandType } from "../types/command_type";

export const testCommand: CommandType = {
  name: "test",
  description: "A test command",
  type: 1,
};

export const allCommands: CommandType[] = [testCommand];
