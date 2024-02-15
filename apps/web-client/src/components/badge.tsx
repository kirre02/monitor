import { Check } from "monitor-sdk/models";

interface StatusBadgeProps {
  status: Check | undefined;
}

function StatusBadge({ status }: StatusBadgeProps) {
  const up = status?.up;
  let badgeColor;
  let badgeText;

  if (up === true) {
    badgeColor = "green";
    badgeText = "Up";
  } else if (up === false) {
    badgeColor = "red";
    badgeText = "Down";
  } else {
    badgeColor = "gray";
    badgeText = "Unknown";
  }

  // @ts-ignore
  return <Badge color={badgeColor}>{badgeText}</Badge>;
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
