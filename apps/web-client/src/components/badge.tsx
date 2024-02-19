import { Status } from "monitor-sdk/models";

interface StatusBadgeProps {
  status: Status | undefined;
}

function StatusBadge({ status }: StatusBadgeProps) {
  const up = status?.up;
  return up ? (
    <Badge color="green">Up</Badge>
  ) : up === false ? (
    <Badge color="red">Down</Badge>
  ) : (
    <Badge color="gray">Unknown</Badge>
  );
}

interface BadgeProps {
  color: "green" | "red" | "gray";
  children?: React.ReactNode;
}

function Badge({ color, children }: BadgeProps) {
  const validColors: Record<"green" | "red" | "gray", [string, string]> = {
    green: ["bg-green-100", "text-green-800"],
    red: ["bg-red-100", "text-red-800"],
    gray: ["bg-gray-100", "text-gray-800"],
  };

  const [bgColor, textColor] = validColors[color] || validColors["gray"];

  return (
    <span
      className={`inline-flex items-center rounded-md px-1.5 py-0.5 text-sm font-medium uppercase ${bgColor} ${textColor}`}
    >
      {children}
    </span>
  );
}

export default StatusBadge;
