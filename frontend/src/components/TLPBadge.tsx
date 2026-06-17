import React from 'react';

interface TLPBadgeProps {
  level: 'white' | 'green' | 'amber' | 'red' | string;
}

export const TLPBadge: React.FC<TLPBadgeProps> = ({ level }) => {
  const normalizedLevel = level.toLowerCase();
  
  let colorClasses = 'bg-gray-800 text-gray-300 border-gray-700';
  let label = 'TLP:WHITE';

  switch (normalizedLevel) {
    case 'green':
      colorClasses = 'bg-green-950/50 text-success border-success/30';
      label = 'TLP:GREEN';
      break;
    case 'amber':
      colorClasses = 'bg-amber-950/50 text-warning border-warning/30';
      label = 'TLP:AMBER';
      break;
    case 'red':
      colorClasses = 'bg-red-950/50 text-danger border-danger/30';
      label = 'TLP:RED';
      break;
    case 'white':
    default:
      colorClasses = 'bg-slate-800/80 text-slate-300 border-slate-700';
      label = 'TLP:WHITE';
      break;
  }

  return (
    <span className={`inline-flex items-center px-2 py-0.5 rounded text-xs font-semibold uppercase tracking-wider border ${colorClasses}`}>
      <span className="w-1.5 h-1.5 rounded-full bg-current mr-1.5 animate-pulse" />
      {label}
    </span>
  );
};
export default TLPBadge;
