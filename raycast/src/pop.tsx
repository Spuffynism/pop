import { closeMainWindow } from "@raycast/api";
import { execa } from "execa";
import { shellEnv } from "shell-env";

interface PopArguments {
  project: string;
  branch?: string;
}

export default async (props: { arguments: PopArguments }) => {
  const { project, branch } = props.arguments;

  await pop(project, branch);
  await closeMainWindow();
};

const pop = async (project: string, branch?: string) => {
  try {
    const { PATH: path } = await shellEnv();
    return await execa(`pop ${project} ${branch}`, { env: { PATH: path }, shell: true });
  } catch (error) {
    console.error(error);
  }
};
