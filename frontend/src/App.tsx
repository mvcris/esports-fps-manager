import { LoadPlayers } from "../wailsjs/go/backend/Game";

function App() {

  const onClick = () => {
    LoadPlayers();
  }

  return (
    <div className="min-h-screen bg-gray-100 flex items-center justify-center p-6">
      <div className="max-w-md w-full rounded-xl bg-white shadow-lg p-6 text-center space-y-4">
        <h1 className="text-2xl font-bold text-gray-900">Tailwind Test</h1>
        <p className="text-gray-600">If this styled card appears, Tailwind is working.</p>
        <div className="flex items-center justify-center gap-3">
          <span className="inline-block h-6 w-6 rounded-full bg-red-500"></span>
          <span className="inline-block h-6 w-6 rounded-full bg-green-500"></span>
          <span className="inline-block h-6 w-6 rounded-full bg-blue-500"></span>
        </div>
        <button className="px-4 py-2 rounded-lg bg-indigo-600 text-white hover:bg-indigo-700 transition" onClick={onClick}>
          Tailwind Button
        </button>
      </div>
    </div>
  )
}

export default App
