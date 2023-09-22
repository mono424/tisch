"use client";
import ArrowUpwardRounded from '@material-ui/icons/ArrowUpwardRounded';
import ArrowDownwardRounded from '@material-ui/icons/ArrowDownwardRounded';

interface ControlsProps {
    onUp: () => void;
    onDown: () => void;
}

export default function Controls({ onUp, onDown }: ControlsProps) {
    return <section className="w-full">
        <div className="flex justify-center">
            <button onClick={onUp} className="px-7 py-3 bg-gray-800 rounded-lg mx-2">
                <ArrowUpwardRounded />
            </button>
            <button onClick={onDown} className="px-7 py-3 bg-gray-800 rounded-lg mx-2">
                <ArrowDownwardRounded />
            </button>
        </div>
    </section>;
}