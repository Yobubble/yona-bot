export interface CommandType {
  name: string;
  description: string;
  type: number;
  options?: any;
  choices?: any;
  required?: boolean;
}
