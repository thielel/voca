export interface AdminResult {
  id: string
  session_id: string
  extraversion: number
  agreeableness: number
  conscientiousness: number
  emotional_stability: number
  openness: number
  created_at: string
}

export const useAdmin = () => {
  const results = ref<AdminResult[]>([])
  const isLoading = ref(false)
  const error = ref<string | null>(null)

  const fetchResults = async () => {
    isLoading.value = true
    error.value = null

    try {
      const data = await $fetch<AdminResult[]>('http://localhost:8080/api/admin/results')
      results.value = data || []
    } catch (e) {
      error.value = 'Fehler beim Laden der Ergebnisse'
      console.error('Failed to fetch results:', e)
    } finally {
      isLoading.value = false
    }
  }

  const formatDate = (dateString: string) => {
    if (!dateString) return '-'
    const date = new Date(dateString)
    return date.toLocaleDateString('de-DE', {
      year: 'numeric',
      month: '2-digit',
      day: '2-digit',
      hour: '2-digit',
      minute: '2-digit'
    })
  }

  const truncateId = (id: string, length = 8) => {
    return id.length > length ? id.slice(0, length) + '...' : id
  }

  return {
    results,
    isLoading,
    error,
    fetchResults,
    formatDate,
    truncateId
  }
}
