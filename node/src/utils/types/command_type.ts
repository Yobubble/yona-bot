export class CommandType {
  constructor (
  public name: string,
  public description: string,
  public type: number,
  public options?: any,
  public choices?: any,
  public required?: boolean,
  public integration_types?: number[],
  public contexts?: number[],
  ) {}
}
