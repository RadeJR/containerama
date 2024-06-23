class ContainerData {
  image: string;
  name?: string;
  env?: string;
  cmd?: string;
  ports?: string;
  volumes?: string;
  entrypoint?: string;
  labels?: string;
  networkDisabled: boolean;

  constructor() {
    this.image = "";
    this.name = "";
    this.env = "";
    this.cmd = "";
    this.ports = "";
    this.volumes = "";
    this.entrypoint = "";
    this.labels = "";
    this.networkDisabled = false;
  }
}

export { ContainerData };
