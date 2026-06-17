import React from 'react';

interface ConfidenceBarProps {
  value: number;
}

export const ConfidenceBar: React.FC<ConfidenceBarProps> = ({ value }) => {
  const percentage = Math.min(Math.max(value, 0), 100);
  
  let barColor = 'bg-danger';
  let textColor = 'text-danger';

  if (percentage >= 70) {
    barColor = 'bg-success';
    textColor = 'text-success';
  } else if (percentage >= 40) {
    barColor = 'bg-warning';
    textColor = 'text-warning';
  }

  return (
    <div className="flex items-center gap-2 w-full max-w-[120px]">
      <div className="w-full h-1.5 bg-border rounded-full overflow-hidden">
        <div 
          className={`h-full rounded-full transition-all duration-300 ${barColor}`} 
          style={{ width: `${percentage}%` }}
        />
      </div>
      <span className={`text-xs font-semibold whitespace-nowrap ${textColor}`}>
        {percentage}%
      </span>
    </div>
  );
};

export default ConfidenceBar;
