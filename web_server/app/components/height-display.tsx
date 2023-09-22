"use client";

interface HeightDisplayProps {
    value: number;
}

export default function HeightDisplay({value}: HeightDisplayProps) {
    return <section className="w-full">
        <div className="flex justify-center text-6xl font-mono px-8 py-4 rounded-xl bg-slate-900">
           {value.toFixed(1)}cm 
        </div>
    </section>;
}