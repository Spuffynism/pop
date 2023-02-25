import { showToast, Toast, closeMainWindow } from "@raycast/api";
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
  let { PATH, LC_ALL } = await shellEnv();

  if (!LC_ALL.includes(".UTF-8")) {
    LC_ALL += ".UTF-8";
  }

  return execa(`pop ${project} ${branch}`, { env: { PATH, LC_ALL }, shell: true, detached: true, cleanup: false });
};
