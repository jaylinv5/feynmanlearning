import { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import { Book, ChevronRight, Filter, Search, Star, Clock } from 'lucide-react';
import { knowledgeApi } from '@/services/knowledge';
import type { KnowledgePointListItem } from '@/types/knowledge';

const subjectMap: Record<string, { name: string; color: string }> = {
  math: { name: '数学', color: 'bg-blue-100 text-blue-700' },
  chinese: { name: '语文', color: 'bg-red-100 text-red-700' },
  english: { name: '英语', color: 'bg-green-100 text-green-700' },
};

const difficultyMap = ['', '简单', '中等', '困难'];

export default function KnowledgeList() {
  const [subject, setSubject] = useState('math');
  const [grade, setGrade] = useState(7);
  const [knowledgeList, setKnowledgeList] = useState<KnowledgePointListItem[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    loadKnowledgeList();
  }, [subject, grade]);

  const loadKnowledgeList = async () => {
    try {
      setLoading(true);
      const data = await knowledgeApi.getBySubjectAndGrade(subject, grade);
      setKnowledgeList(data);
    } catch (error) {
      console.error('加载知识点列表失败:', error);
    } finally {
      setLoading(false);
    }
  };

  // 按章节分组
  const chapters = knowledgeList.reduce((acc, item) => {
    if (!acc[item.chapter]) {
      acc[item.chapter] = [];
    }
    acc[item.chapter].push(item);
    return acc;
  }, {} as Record<string, KnowledgePointListItem[]>);

  if (loading) {
    return (
      <div className="min-h-screen bg-gray-50 flex items-center justify-center">
        <div className="text-center">
          <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-primary-600 mx-auto"></div>
          <p className="mt-4 text-gray-600">加载中...</p>
        </div>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-gray-50">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        {/* 头部 */}
        <div className="mb-8">
          <h1 className="text-3xl font-bold text-gray-900 mb-2">知识点学习</h1>
          <p className="text-gray-600">选择你要学习的知识点，开启费曼学习之旅</p>
        </div>

        {/* 筛选栏 */}
        <div className="bg-white rounded-lg shadow-sm p-4 mb-6">
          <div className="flex flex-wrap gap-4 items-center">
            <div className="flex items-center gap-2">
              <Filter className="w-5 h-5 text-gray-500" />
              <span className="text-gray-700 font-medium">学科:</span>
              <select
                value={subject}
                onChange={(e) => setSubject(e.target.value)}
                className="border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-primary-500"
              >
                <option value="math">数学</option>
                <option value="chinese">语文</option>
                <option value="english">英语</option>
              </select>
            </div>

            <div className="flex items-center gap-2">
              <span className="text-gray-700 font-medium">年级:</span>
              <select
                value={grade}
                onChange={(e) => setGrade(Number(e.target.value))}
                className="border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-primary-500"
              >
                <option value={7}>七年级</option>
                <option value={8}>八年级</option>
                <option value={9}>九年级</option>
              </select>
            </div>
          </div>
        </div>

        {/* 知识点列表 */}
        <div className="space-y-6">
          {Object.entries(chapters).map(([chapterName, items]) => (
            <div key={chapterName} className="bg-white rounded-lg shadow-sm overflow-hidden">
              <div className="bg-gradient-to-r from-primary-50 to-primary-100 px-6 py-4 border-b border-primary-200">
                <h2 className="text-xl font-semibold text-primary-900 flex items-center gap-2">
                  <Book className="w-5 h-5" />
                  {chapterName}
                </h2>
                <p className="text-sm text-primary-700 mt-1">
                  共 {items.length} 个知识点
                </p>
              </div>

              <div className="divide-y divide-gray-200">
                {items.map((item) => (
                  <Link
                    key={item.id}
                    to={`/knowledge/${item.id}`}
                    className="px-6 py-4 hover:bg-gray-50 transition-colors flex items-center justify-between group"
                  >
                    <div className="flex-1">
                      <div className="flex items-center gap-3 mb-2">
                        <h3 className="text-lg font-medium text-gray-900 group-hover:text-primary-600 transition-colors">
                          {item.name}
                        </h3>
                        <span className={`px-2 py-1 rounded-full text-xs font-medium ${subjectMap[item.subject].color}`}>
                          {subjectMap[item.subject].name}
                        </span>
                        <span className="px-2 py-1 bg-gray-100 text-gray-700 rounded-full text-xs font-medium">
                          {difficultyMap[item.difficulty]}
                        </span>
                        {item.mastered && (
                          <span className="px-2 py-1 bg-success-100 text-success-700 rounded-full text-xs font-medium">
                            已掌握
                          </span>
                        )}
                      </div>

                      <div className="flex items-center gap-6 text-sm text-gray-500">
                        <div className="flex items-center gap-1">
                          <Clock className="w-4 h-4" />
                          <span>预计 {item.estimatedTime} 分钟</span>
                        </div>
                        {item.tags && (
                          <div className="flex items-center gap-2">
                            {item.tags.split(',').map((tag) => (
                              <span key={tag} className="text-gray-600">
                                #{tag}
                              </span>
                            ))}
                          </div>
                        )}
                      </div>
                    </div>

                    <ChevronRight className="w-5 h-5 text-gray-400 group-hover:text-primary-600 transition-colors" />
                  </Link>
                ))}
              </div>
            </div>
          ))}
        </div>

        {Object.keys(chapters).length === 0 && (
          <div className="bg-white rounded-lg shadow-sm p-12 text-center">
            <Book className="w-16 h-16 text-gray-300 mx-auto mb-4" />
            <h3 className="text-lg font-medium text-gray-900 mb-2">暂无知识点</h3>
            <p className="text-gray-500">该年级学科下还没有添加任何知识点</p>
          </div>
        )}
      </div>
    </div>
  );
}
