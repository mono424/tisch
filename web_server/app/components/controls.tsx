"use client";
import ArrowUpwardRounded from '@material-ui/icons/ArrowUpwardRounded';
import ArrowDownwardRounded from '@material-ui/icons/ArrowDownwardRounded';

export default function Controls() {
    return <section className="w-full">
        <div className="flex justify-center">
            <button className="px-7 py-3 bg-gray-800 rounded-lg mx-2">
                <ArrowUpwardRounded />
            </button>
            <button className="px-7 py-3 bg-gray-800 rounded-lg mx-2">
                <ArrowDownwardRounded />
            </button>
        </div>
    </section>;
}