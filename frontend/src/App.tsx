import { Routes, Route } from 'react-router-dom'
import { Toaster } from 'sonner'
import KnowledgeList from './pages/knowledge/List'
import './App.css'

function App() {
  return (
    <div className="min-h-screen bg-gray-50">
      <Routes>
        <Route path="/" element={<KnowledgeList />} />
        <Route path="/knowledge/:id" element={<div>知识点详情页开发中...</div>} />
      </Routes>
      <Toaster position="top-right" />
    </div>
  )
}

export default App
