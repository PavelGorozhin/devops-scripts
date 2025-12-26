// types.ts

// Define types for various data structures and functionality
import { IEnvironmentStrategy, IEnvironmentOption } from './environment';

export interface IConfig {
  env: IEnvironmentOption;
  deploy: DeployConfig;
  logs: LogsConfig;
}

export interface IEnvironmentStrategy {
  name: string;
  options: IEnvironmentOption[];
}

export interface IEnvironmentOption {
  name: string;
  description: string;
  envVariables: { [key: string]: string };
}

export interface DeployConfig {
  type: 'aws' | 'gcp';
  account: string;
  region: string;
}

export interface LogsConfig {
  type: 'console' | 'file';
  file: { filename: string; maxLines: number };
}

export interface IAppConfig {
  config: IConfig;
}