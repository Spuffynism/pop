import { showToast, Toast, closeMainWindow, confirmAlert } from "@raycast/api";
import { execa } from "execa";
import { shellEnv } from "shell-env";

interface PopArguments {
  project: string;
  branch?: string;
}

export default async (props: { arguments: PopArguments }) => {
  const { project, branch } = props.arguments;

  try {
    await pop(project, branch);
    await closeMainWindow();
  } catch (error: any) {
    await showToast({
      style: Toast.Style.Failure,
      title: `Couldn't open '${project}'!`,
      message: error.stderr || "Unknown error",
    });
  }
};

const pop = async (project: string, branch?: string) => {
  const { PATH: path } = await shellEnv();
  return execa(`pop ${project} ${branch}`, { env: { PATH: path }, shell: true, detached: true, cleanup: false });
};
