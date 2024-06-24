export class StackFileData {
  name: string;
  content: string;

  constructor() {
    this.name = "";
    this.content = "";
  }
}

export class Stack {
  Id: number;
  CreatedAt: Date;
  UpdatedAt: Date;
  Name: string;
  PathToFile: string;
  Webhook: string | null;
  UserId: number;
  RepositoryId: number | null;
  Branch: string | null;

  constructor(
    id: number,
    createdAt: Date,
    updatedAt: Date,
    name: string,
    pathToFile: string,
    webhook: string | null,
    userId: number,
    repositoryId: number | null,
    branch: string | null
  ) {
    this.Id = id;
    this.CreatedAt = createdAt;
    this.UpdatedAt = updatedAt;
    this.Name = name;
    this.PathToFile = pathToFile;
    this.Webhook = webhook;
    this.UserId = userId;
    this.RepositoryId = repositoryId;
    this.Branch = branch;
  }
}
