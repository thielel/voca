/**
 * Design Philosophy: Neo-Brutalism meets Swiss Modernism
 * - Bold, unapologetic typography with strong hierarchy
 * - High contrast color blocks without gradients
 * - Geometric precision with asymmetric layouts
 * - Raw, honest presentation of information
 */

import { useState } from "react";
import { Card } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import { 
  Sparkles, 
  Target, 
  Heart, 
  GraduationCap,
  Check,
  Star
} from "lucide-react";

interface NameOption {
  name: string;
  description: string;
  appeal: string;
}

interface Category {
  id: string;
  title: string;
  rationale: string;
  color: string;
  icon: React.ReactNode;
  names: NameOption[];
}

const categories: Category[] = [
  {
    id: "modern",
    title: "Modern & Tech-Forward",
    rationale: "This category leverages the appeal of technology and innovation, which resonates strongly with Gen Z. The names are short, memorable, and hint at a data-driven, intelligent system. They position the test as a cutting-edge tool for the future of work.",
    color: "oklch(0.6 0.25 250)",
    icon: <Sparkles className="w-6 h-6" />,
    names: [
      {
        name: "CareerAI",
        description: "A straightforward and descriptive name that clearly communicates the test's purpose and AI integration.",
        appeal: "Simple, easy to remember, and directly states the value proposition."
      },
      {
        name: "FutureFit",
        description: "Implies a forward-looking approach to career planning, helping users find a career that fits them.",
        appeal: "Empowering and aspirational, suggesting a proactive approach to their future."
      },
      {
        name: "TalentSync",
        description: "Suggests a synchronization between the user's talents and potential career paths.",
        appeal: "Modern and dynamic, implying a personalized and accurate matching process."
      },
      {
        name: "Vocate",
        description: "A blend of 'vocation' and 'advocate,' suggesting a tool that champions the user's career journey.",
        appeal: "Short, unique, and memorable, with a sophisticated yet approachable feel."
      }
    ]
  },
  {
    id: "empowering",
    title: "Empowering & Aspirational",
    rationale: "These names focus on the user's personal journey of self-discovery and empowerment. They are designed to be inspiring and motivational, encouraging young people to explore their potential and take control of their future. The language is positive and aspirational.",
    color: "oklch(0.7 0.2 40)",
    icon: <Target className="w-6 h-6" />,
    names: [
      {
        name: "MyPath",
        description: "A personal and empowering name that emphasizes ownership of one's career journey.",
        appeal: "Simple, personal, and empowering, giving users a sense of agency."
      },
      {
        name: "CareerSpark",
        description: "Suggests that the test can ignite a passion or spark an idea for a future career.",
        appeal: "Creative and evocative, implying a moment of inspiration and discovery."
      },
      {
        name: "The Compass",
        description: "A classic metaphor for guidance and direction, positioning the test as a reliable tool for navigation.",
        appeal: "Trustworthy and reliable, offering a sense of security and clarity in a complex world."
      },
      {
        name: "Horizons",
        description: "Evokes a sense of exploration and boundless possibilities, encouraging users to look toward the future.",
        appeal: "Aspirational and optimistic, appealing to the desire for adventure and new experiences."
      }
    ]
  },
  {
    id: "friendly",
    title: "Friendly & Approachable",
    rationale: "This category uses friendly and conversational language to make the test feel less like a formal assessment and more like a helpful guide. The names are designed to be welcoming and non-intimidating, fostering a sense of trust and companionship.",
    color: "oklch(0.85 0.18 95)",
    icon: <Heart className="w-6 h-6" />,
    names: [
      {
        name: "CareerPal",
        description: "A friendly and approachable name that positions the test as a helpful companion.",
        appeal: "Warm and inviting, suggesting a supportive and non-judgmental experience."
      },
      {
        name: "JobCompass",
        description: "A combination of a familiar term ('job') with a metaphor for guidance ('compass').",
        appeal: "Clear and practical, with a friendly and helpful tone."
      },
      {
        name: "Guidely",
        description: "A playful and modern-sounding name that suggests guidance and support.",
        appeal: "Unique and memorable, with a friendly and approachable feel."
      },
      {
        name: "Voca",
        description: "A short, catchy, and friendly-sounding name derived from 'vocation.'",
        appeal: "Easy to remember and pronounce, with a modern and minimalist appeal."
      }
    ]
  },
  {
    id: "scientific",
    title: "Scientific & Credible",
    rationale: "While still being youth-friendly, these names subtly hint at the scientific foundation of the test (BFI-K). They are designed to build trust and credibility by suggesting a rigorous and evidence-based approach. The names are more formal but still accessible.",
    color: "oklch(0.5 0.2 150)",
    icon: <GraduationCap className="w-6 h-6" />,
    names: [
      {
        name: "PersonaProfile",
        description: "A professional-sounding name that suggests a detailed and comprehensive personality assessment.",
        appeal: "Sophisticated and credible, appealing to users who value accuracy and depth."
      },
      {
        name: "MindsetMap",
        description: "A metaphorical name that suggests a mapping of the user's personality and cognitive style.",
        appeal: "Intriguing and thought-provoking, implying a deeper level of self-understanding."
      },
      {
        name: "The Core Five",
        description: "A direct reference to the Big Five personality traits, positioning the test as a credible instrument.",
        appeal: "Clear and informative, for users who are interested in the science behind the test."
      },
      {
        name: "Cognify",
        description: "A modern and scientific-sounding name that suggests a process of understanding and self-awareness.",
        appeal: "Short, memorable, and tech-forward, with a hint of scientific rigor."
      }
    ]
  }
];

export default function Home() {
  const [selectedNames, setSelectedNames] = useState<Set<string>>(new Set());
  const [activeCategory, setActiveCategory] = useState("modern");

  const toggleNameSelection = (name: string) => {
    const newSelection = new Set(selectedNames);
    if (newSelection.has(name)) {
      newSelection.delete(name);
    } else {
      newSelection.add(name);
    }
    setSelectedNames(newSelection);
  };

  return (
    <div className="min-h-screen bg-background">
      {/* Hero Section - Asymmetric layout with geometric background */}
      <section className="relative overflow-hidden border-b-[6px] border-black">
        <div className="absolute inset-0 opacity-10">
          <img 
            src="/images/hero-geometric.jpg" 
            alt="" 
            className="w-full h-full object-cover"
          />
        </div>
        <div className="container relative py-16 md:py-24">
          <div className="grid md:grid-cols-12 gap-8 items-center">
            <div className="md:col-span-7">
              <Badge 
                className="mb-6 text-base px-4 py-2 brutalist-border bg-[oklch(0.85_0.18_95)] text-black font-bold"
              >
                AI-POWERED NAMING
              </Badge>
              <h1 className="text-5xl md:text-7xl font-black mb-6 leading-tight">
                Career Test
                <br />
                <span className="text-[oklch(0.6_0.25_250)]">Name Explorer</span>
              </h1>
              <p className="text-xl md:text-2xl mb-8 font-medium max-w-2xl">
                16 creative naming options for your AI-infused personality test targeting 15-18 year olds. Based on BFI-K KJ with future voice capabilities.
              </p>
              <div className="flex flex-wrap gap-4">
                <Button 
                  size="lg" 
                  className="brutalist-border brutalist-shadow bg-[oklch(0.6_0.25_250)] hover:bg-[oklch(0.5_0.25_250)] text-white font-bold text-lg px-8 py-6 transition-all hover:translate-x-1 hover:translate-y-1 hover:shadow-none"
                  onClick={() => document.getElementById('categories')?.scrollIntoView({ behavior: 'smooth' })}
                >
                  EXPLORE NAMES
                </Button>
                <Button 
                  size="lg" 
                  variant="outline"
                  className="brutalist-border bg-white hover:bg-[oklch(0.98_0.005_100)] font-bold text-lg px-8 py-6"
                  onClick={() => {
                    const selected = Array.from(selectedNames);
                    if (selected.length > 0) {
                      alert(`Your selected names:\n\n${selected.join('\n')}`);
                    } else {
                      alert('Select some names first by clicking on them!');
                    }
                  }}
                >
                  VIEW SELECTED ({selectedNames.size})
                </Button>
              </div>
            </div>
            <div className="md:col-span-5 flex justify-end">
              <div className="brutalist-border-thick brutalist-shadow bg-white p-8 max-w-sm">
                <div className="flex items-center gap-3 mb-4">
                  <Star className="w-8 h-8 text-[oklch(0.7_0.2_40)]" fill="currentColor" />
                  <h3 className="text-2xl font-black">Quick Stats</h3>
                </div>
                <div className="space-y-4">
                  <div className="border-t-4 border-black pt-4">
                    <div className="text-4xl font-black text-[oklch(0.6_0.25_250)]">16</div>
                    <div className="text-sm font-bold">UNIQUE NAMES</div>
                  </div>
                  <div className="border-t-4 border-black pt-4">
                    <div className="text-4xl font-black text-[oklch(0.7_0.2_40)]">4</div>
                    <div className="text-sm font-bold">CATEGORIES</div>
                  </div>
                  <div className="border-t-4 border-black pt-4">
                    <div className="text-4xl font-black text-[oklch(0.85_0.18_95)]">{selectedNames.size}</div>
                    <div className="text-sm font-bold">SELECTED</div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </section>

      {/* Categories Section */}
      <section id="categories" className="py-16 md:py-24">
        <div className="container">
          <div className="mb-12">
            <h2 className="text-4xl md:text-5xl font-black mb-4">
              NAMING CATEGORIES
            </h2>
            <p className="text-lg md:text-xl max-w-3xl">
              Each category represents a distinct branding approach. Click on any name to select it for your shortlist.
            </p>
          </div>

          <Tabs value={activeCategory} onValueChange={setActiveCategory} className="w-full">
            <TabsList className="grid w-full grid-cols-2 md:grid-cols-4 gap-4 bg-transparent h-auto p-0 mb-8">
              {categories.map((category) => (
                <TabsTrigger
                  key={category.id}
                  value={category.id}
                  className="brutalist-border data-[state=active]:brutalist-shadow bg-white data-[state=active]:bg-[oklch(0.15_0_0)] data-[state=active]:text-white font-bold text-sm md:text-base py-4 px-3 transition-all data-[state=active]:translate-x-1 data-[state=active]:translate-y-1"
                  style={{
                    borderColor: category.color,
                    borderWidth: '4px'
                  }}
                >
                  <span className="flex items-center gap-2 justify-center">
                    {category.icon}
                    <span className="hidden md:inline">{category.title}</span>
                    <span className="md:hidden">{category.title.split(' ')[0]}</span>
                  </span>
                </TabsTrigger>
              ))}
            </TabsList>

            {categories.map((category) => (
              <TabsContent key={category.id} value={category.id} className="mt-0">
                {/* Category Header */}
                <div 
                  className="brutalist-border-thick p-8 mb-8"
                  style={{ 
                    backgroundColor: category.color,
                    color: 'white'
                  }}
                >
                  <div className="flex items-start gap-4 mb-4">
                    <div className="p-3 bg-white text-black brutalist-border">
                      {category.icon}
                    </div>
                    <div className="flex-1">
                      <h3 className="text-3xl md:text-4xl font-black mb-2">
                        {category.title}
                      </h3>
                    </div>
                  </div>
                  <p className="text-lg md:text-xl font-medium leading-relaxed">
                    {category.rationale}
                  </p>
                </div>

                {/* Name Cards Grid */}
                <div className="grid md:grid-cols-2 gap-6">
                  {category.names.map((nameOption) => {
                    const isSelected = selectedNames.has(nameOption.name);
                    return (
                      <Card
                        key={nameOption.name}
                        className={`brutalist-border bg-white p-6 cursor-pointer transition-all hover:translate-x-1 hover:translate-y-1 ${
                          isSelected ? 'brutalist-shadow' : ''
                        }`}
                        style={{
                          borderColor: isSelected ? category.color : 'oklch(0.15 0 0)',
                          borderWidth: isSelected ? '6px' : '4px'
                        }}
                        onClick={() => toggleNameSelection(nameOption.name)}
                      >
                        <div className="flex items-start justify-between mb-4">
                          <h4 className="text-2xl md:text-3xl font-black">
                            {nameOption.name}
                          </h4>
                          {isSelected && (
                            <div 
                              className="p-2 brutalist-border"
                              style={{ backgroundColor: category.color }}
                            >
                              <Check className="w-5 h-5 text-white" />
                            </div>
                          )}
                        </div>
                        <p className="text-base font-medium mb-4 leading-relaxed">
                          {nameOption.description}
                        </p>
                        <div className="border-t-4 border-black pt-4">
                          <div className="text-xs font-bold uppercase mb-2 opacity-60">
                            APPEAL TO YOUTH
                          </div>
                          <p className="text-sm font-medium">
                            {nameOption.appeal}
                          </p>
                        </div>
                      </Card>
                    );
                  })}
                </div>
              </TabsContent>
            ))}
          </Tabs>
        </div>
      </section>

      {/* Footer */}
      <footer className="border-t-[6px] border-black py-12 bg-[oklch(0.15_0_0)] text-white">
        <div className="container">
          <div className="grid md:grid-cols-2 gap-8">
            <div>
              <h3 className="text-2xl font-black mb-4">ABOUT THIS PROJECT</h3>
              <p className="font-medium leading-relaxed">
                This interactive naming explorer presents 16 creative options for an AI-infused personality test 
                targeting 15-18 year olds. The test is based on the scientifically validated BFI-K KJ framework 
                and will feature voice capabilities and LLM chat functionality.
              </p>
            </div>
            <div>
              <h3 className="text-2xl font-black mb-4">DESIGN APPROACH</h3>
              <p className="font-medium leading-relaxed">
                Built with a Neo-Brutalist design philosophy featuring bold typography, high-contrast color blocks, 
                and asymmetric layouts. Each category has its own distinct color for instant visual differentiation.
              </p>
            </div>
          </div>
          <div className="mt-8 pt-8 border-t-4 border-white text-center">
            <p className="font-bold">
              Created with Manus AI • 2026
            </p>
          </div>
        </div>
      </footer>
    </div>
  );
}
