import type { ReactNode } from 'react';

type TooltipProps = {
  label: string;
  children: ReactNode;
};

// Lightweight CSS tooltip shown on hover. Themed via the design tokens
// (inverted for contrast).
export function Tooltip({ label, children }: TooltipProps) {
  return (
    <span className="group relative inline-flex">
      {children}
      <span
        role="tooltip"
        className="pointer-events-none absolute -top-9 left-1/2 -translate-x-1/2 whitespace-nowrap rounded bg-foreground px-2 py-1 text-xs text-background opacity-0 transition-opacity duration-150 group-hover:opacity-100"
      >
        {label}
      </span>
    </span>
  );
}
