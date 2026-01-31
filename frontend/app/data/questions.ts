/**
 * IPIP Big Five Personality Test - 50 Item Version
 * Source: International Personality Item Pool (https://ipip.ori.org/)
 * License: Public Domain
 *
 * German translations of the official IPIP 50-item scale
 */

export type Trait =
  | 'extraversion'
  | 'agreeableness'
  | 'conscientiousness'
  | 'emotionalStability'
  | 'openness'

export interface Question {
  id: number
  text: string
  trait: Trait
  reversed: boolean
}

export interface Answer {
  questionId: number
  value: number // 1-5 Likert scale
}

export interface PersonalityResult {
  id?: string
  extraversion: number
  agreeableness: number
  conscientiousness: number
  emotionalStability: number
  openness: number
  createdAt?: string
}

/**
 * Likert scale options (German)
 * 1: Trifft überhaupt nicht zu (Very Inaccurate)
 * 2: Trifft eher nicht zu (Moderately Inaccurate)
 * 3: Weder zutreffend noch unzutreffend (Neither Accurate Nor Inaccurate)
 * 4: Trifft eher zu (Moderately Accurate)
 * 5: Trifft voll und ganz zu (Very Accurate)
 */
export const likertOptions = [
  { value: 1, label: 'Trifft überhaupt nicht zu' },
  { value: 2, label: 'Trifft eher nicht zu' },
  { value: 3, label: 'Weder zutreffend noch unzutreffend' },
  { value: 4, label: 'Trifft eher zu' },
  { value: 5, label: 'Trifft voll und ganz zu' },
] as const

/**
 * IPIP 50-Item Big Five Questionnaire
 * Items alternate between traits and polarity for validity
 */
export const questions: Question[] = [
  // 1-10: First round of all 5 traits (alternating polarity)
  { id: 1, text: 'Ich bin der Mittelpunkt jeder Party.', trait: 'extraversion', reversed: false },
  { id: 2, text: 'Ich interessiere mich wenig für andere.', trait: 'agreeableness', reversed: true },
  { id: 3, text: 'Ich bin immer gut vorbereitet.', trait: 'conscientiousness', reversed: false },
  { id: 4, text: 'Ich bin leicht gestresst.', trait: 'emotionalStability', reversed: true },
  { id: 5, text: 'Ich habe einen reichen Wortschatz.', trait: 'openness', reversed: false },
  { id: 6, text: 'Ich rede nicht viel.', trait: 'extraversion', reversed: true },
  { id: 7, text: 'Ich interessiere mich für andere Menschen.', trait: 'agreeableness', reversed: false },
  { id: 8, text: 'Ich lasse meine Sachen herumliegen.', trait: 'conscientiousness', reversed: true },
  { id: 9, text: 'Ich bin die meiste Zeit entspannt.', trait: 'emotionalStability', reversed: false },
  { id: 10, text: 'Ich habe Schwierigkeiten, abstrakte Ideen zu verstehen.', trait: 'openness', reversed: true },

  // 11-20: Second round
  { id: 11, text: 'Ich fühle mich wohl unter Menschen.', trait: 'extraversion', reversed: false },
  { id: 12, text: 'Ich beleidige andere.', trait: 'agreeableness', reversed: true },
  { id: 13, text: 'Ich achte auf Details.', trait: 'conscientiousness', reversed: false },
  { id: 14, text: 'Ich mache mir Sorgen über Dinge.', trait: 'emotionalStability', reversed: true },
  { id: 15, text: 'Ich habe eine lebhafte Fantasie.', trait: 'openness', reversed: false },
  { id: 16, text: 'Ich halte mich im Hintergrund.', trait: 'extraversion', reversed: true },
  { id: 17, text: 'Ich fühle mit den Gefühlen anderer mit.', trait: 'agreeableness', reversed: false },
  { id: 18, text: 'Ich bringe Dinge durcheinander.', trait: 'conscientiousness', reversed: true },
  { id: 19, text: 'Ich bin selten niedergeschlagen.', trait: 'emotionalStability', reversed: false },
  { id: 20, text: 'Ich interessiere mich nicht für abstrakte Ideen.', trait: 'openness', reversed: true },

  // 21-30: Third round
  { id: 21, text: 'Ich beginne Gespräche.', trait: 'extraversion', reversed: false },
  { id: 22, text: 'Ich interessiere mich nicht für die Probleme anderer.', trait: 'agreeableness', reversed: true },
  { id: 23, text: 'Ich erledige Aufgaben sofort.', trait: 'conscientiousness', reversed: false },
  { id: 24, text: 'Ich bin leicht aus der Ruhe zu bringen.', trait: 'emotionalStability', reversed: true },
  { id: 25, text: 'Ich habe ausgezeichnete Ideen.', trait: 'openness', reversed: false },
  { id: 26, text: 'Ich habe wenig zu sagen.', trait: 'extraversion', reversed: true },
  { id: 27, text: 'Ich habe ein weiches Herz.', trait: 'agreeableness', reversed: false },
  { id: 28, text: 'Ich vergesse oft, Dinge an ihren Platz zurückzulegen.', trait: 'conscientiousness', reversed: true },
  { id: 29, text: 'Ich rege mich leicht auf.', trait: 'emotionalStability', reversed: true },
  { id: 30, text: 'Ich habe keine gute Vorstellungskraft.', trait: 'openness', reversed: true },

  // 31-40: Fourth round
  { id: 31, text: 'Ich unterhalte mich auf Partys mit vielen verschiedenen Leuten.', trait: 'extraversion', reversed: false },
  { id: 32, text: 'Ich interessiere mich nicht wirklich für andere.', trait: 'agreeableness', reversed: true },
  { id: 33, text: 'Ich mag Ordnung.', trait: 'conscientiousness', reversed: false },
  { id: 34, text: 'Meine Stimmung ändert sich oft.', trait: 'emotionalStability', reversed: true },
  { id: 35, text: 'Ich verstehe Dinge schnell.', trait: 'openness', reversed: false },
  { id: 36, text: 'Ich ziehe es vor, nicht im Mittelpunkt zu stehen.', trait: 'extraversion', reversed: true },
  { id: 37, text: 'Ich nehme mir Zeit für andere.', trait: 'agreeableness', reversed: false },
  { id: 38, text: 'Ich drücke mich vor meinen Pflichten.', trait: 'conscientiousness', reversed: true },
  { id: 39, text: 'Ich habe häufige Stimmungsschwankungen.', trait: 'emotionalStability', reversed: true },
  { id: 40, text: 'Ich benutze schwierige Wörter.', trait: 'openness', reversed: false },

  // 41-50: Fifth round
  { id: 41, text: 'Es macht mir nichts aus, im Mittelpunkt der Aufmerksamkeit zu stehen.', trait: 'extraversion', reversed: false },
  { id: 42, text: 'Ich spüre die Emotionen anderer.', trait: 'agreeableness', reversed: false },
  { id: 43, text: 'Ich halte mich an einen Zeitplan.', trait: 'conscientiousness', reversed: false },
  { id: 44, text: 'Ich werde leicht gereizt.', trait: 'emotionalStability', reversed: true },
  { id: 45, text: 'Ich verbringe Zeit damit, über Dinge nachzudenken.', trait: 'openness', reversed: false },
  { id: 46, text: 'Ich bin still unter Fremden.', trait: 'extraversion', reversed: true },
  { id: 47, text: 'Ich sorge dafür, dass sich andere wohlfühlen.', trait: 'agreeableness', reversed: false },
  { id: 48, text: 'Ich bin genau in meiner Arbeit.', trait: 'conscientiousness', reversed: false },
  { id: 49, text: 'Ich fühle mich oft niedergeschlagen.', trait: 'emotionalStability', reversed: true },
  { id: 50, text: 'Ich bin voller Ideen.', trait: 'openness', reversed: false },
]

/**
 * Get questions filtered by trait
 */
export function getQuestionsByTrait(trait: Trait): Question[] {
  return questions.filter(q => q.trait === trait)
}

/**
 * Get all available traits
 */
export const traits: Trait[] = [
  'extraversion',
  'agreeableness',
  'conscientiousness',
  'emotionalStability',
  'openness',
]

/**
 * Trait display names in German
 */
export const traitLabels: Record<Trait, string> = {
  extraversion: 'Extraversion',
  agreeableness: 'Verträglichkeit',
  conscientiousness: 'Gewissenhaftigkeit',
  emotionalStability: 'Emotionale Stabilität',
  openness: 'Offenheit',
}
