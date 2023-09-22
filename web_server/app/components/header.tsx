export default function Header() {
    return <section className="w-full">
        <div className="container px-4 md:px-6">
            <div className="grid gap-6 items-center">
                <div className="flex flex-col justify-start space-y-8 text-center">
                    <div className="space-y-2">
                        <h1 className="text-3xl font-bold tracking-tighter sm:text-5xl xl:text-6xl/none bg-clip-text text-transparent bg-gradient-to-r from-white to-gray-500">
                        TISCH
                        </h1>
                        <p className="max-w-[600px] text-zinc-200 md:text-xl dark:text-zinc-100 mx-auto">
                        Tisch, make your desk a little more smart. Try it out and control your desk with the buttons below.
                        </p>
                    </div>
                </div>
            </div>
        </div>
    </section>;
}