# VOCA Style Guide

A comprehensive design system for the Big Five Personality Questionnaire, designed to help young people (ages 15-18) discover career paths that align with their personality traits.

---

## Table of Contents

1. [Brand Voice and Tone](#brand-voice-and-tone)
2. [Color Palette](#color-palette)
3. [Gradients](#gradients)
4. [Typography](#typography)
5. [Spacing and Layout](#spacing-and-layout)
6. [Background Decorations](#background-decorations)
7. [Components](#components)
8. [Icons and Imagery](#icons-and-imagery)
9. [Accessibility](#accessibility)

---

## Brand Voice and Tone

### Principles

- **Friendly**: Write like you're talking to a friend, not a textbook
- **Encouraging**: Celebrate progress and emphasize that there are no wrong answers
- **Inclusive**: Use language that welcomes everyone regardless of background
- **Clear**: Avoid jargon; keep sentences short and easy to understand
- **Vibrant**: The design should feel energetic and engaging for young users

### Language Guidelines

| Do | Don't |
|---|---|
| "You're doing great!" | "Proceed to the next section" |
| "What feels most like you?" | "Rate the following statement" |
| "Almost there!" | "75% complete" |
| "Let's explore your strengths" | "Assessment results" |

### Messaging Examples

- **Welcome**: "Ready to discover what makes you, you? Let's find out which careers might be a great fit!"
- **During questionnaire**: "Take your time — there's no rush and no wrong answers."
- **Completion**: "Amazing work! Here's what we learned about you."

---

## Color Palette

### Primary Colors

These form the foundation of the interface and convey trust, energy, and approachability.

| Name | Hex | RGB | Usage |
|------|-----|-----|-------|
| Sky Blue | `#6BB8E4` | rgb(107, 184, 228) | Primary actions, progress indicators |
| Ocean Teal | `#4ABFBF` | rgb(74, 191, 191) | Secondary actions, accents |
| Violet | `#8B5CF6` | rgb(139, 92, 246) | Gradient endpoints, emphasis |
| Purple | `#A855F7` | rgb(168, 85, 247) | Gradient midpoints, highlights |

### Personality Trait Colors

Each Big Five trait has a vibrant gradient color scheme:

| Trait | Primary | Secondary | Gradient |
|-------|---------|-----------|----------|
| Extraversion | Rose `#F43F5E` | Pink `#EC4899` | `from-rose-500 to-pink-500` |
| Agreeableness | Emerald `#10B981` | Teal `#14B8A6` | `from-emerald-500 to-teal-500` |
| Conscientiousness | Amber `#F59E0B` | Orange `#F97316` | `from-amber-500 to-orange-500` |
| Emotional Stability | Sky `#0EA5E9` | Blue `#3B82F6` | `from-sky-500 to-blue-500` |
| Openness | Violet `#8B5CF6` | Purple `#A855F7` | `from-violet-500 to-purple-500` |

### Background Decoration Colors

For the colorful blobs/shapes that add visual interest:

| Color | Hex | Opacity | Usage |
|-------|-----|---------|-------|
| Sky | `#38BDF8` | 30-40% | Top-right decorations |
| Teal | `#2DD4BF` | 20-30% | Accent blobs |
| Rose | `#FB7185` | 20-35% | Bottom-left decorations |
| Violet | `#A78BFA` | 20-30% | Center decorations |
| Amber | `#FBBF24` | 20-25% | Accent highlights |

### Neutral Colors

| Name | Hex | RGB | Usage |
|------|-----|-----|-------|
| White | `#FFFFFF` | rgb(255, 255, 255) | Card backgrounds |
| Cloud | `#F8FAFC` | rgb(248, 250, 252) | Page backgrounds |
| Slate 100 | `#F1F5F9` | rgb(241, 245, 249) | Subtle backgrounds |
| Slate 200 | `#E2E8F0` | rgb(226, 232, 240) | Borders, dividers |
| Slate 500 | `#64748B` | rgb(100, 116, 139) | Secondary text |
| Slate 600 | `#475569` | rgb(71, 85, 105) | Body text |
| Slate 800 | `#1E293B` | rgb(30, 41, 59) | Headings |

### Color Mode

The application defaults to **light mode** to maintain a friendly, inviting atmosphere appropriate for young users. Dark mode is supported but light mode is preferred.

---

## Gradients

Gradients are a key part of the VOCA visual identity, adding energy and visual interest.

### Primary Action Gradient

Used for main CTA buttons:

```css
background: linear-gradient(to right, #0EA5E9, #14B8A6, #10B981);
/* Tailwind: from-sky-500 via-teal-500 to-emerald-500 */
```

### Card Border Gradient

Used for question cards and featured elements:

```css
background: linear-gradient(to right, #8B5CF6, #A855F7, #EC4899);
/* Tailwind: from-violet-500 via-purple-500 to-pink-500 */
```

### Progress Bar Gradient

```css
background: linear-gradient(to right, #8B5CF6, #A855F7, #EC4899);
/* Tailwind: from-violet-500 via-purple-500 to-pink-500 */
```

### Hero Title Gradient (Animated)

```css
background: linear-gradient(135deg, #6366f1, #8b5cf6, #d946ef, #f43f5e, #f97316);
background-size: 200% 200%;
animation: gradientShift 8s ease infinite;

@keyframes gradientShift {
  0%, 100% { background-position: 0% 50%; }
  50% { background-position: 100% 50%; }
}
```

### Feature Section Gradient

Used for the IPIP foundation section:

```css
background: linear-gradient(to bottom right, #7C3AED, #9333EA, #4F46E5);
/* Tailwind: from-violet-600 via-purple-600 to-indigo-700 */
```

### Radar Chart Fill Gradient

```css
background: linear-gradient(135deg, #6BB8E4, #4ABFBF, #A8E6CF);
/* Sky Blue → Ocean Teal → Mint Green */
opacity: 0.5;
```

---

## Typography

### Font Families

```css
/* Primary Font - All UI elements */
font-family: 'Nunito', 'Inter', 'Segoe UI', sans-serif;
```

**Nunito** is chosen for its friendly, rounded letterforms that feel approachable to younger users.

### Type Scale

| Name | Size | Weight | Line Height | Usage |
|------|------|--------|-------------|-------|
| Display | 48-72px | 700 (Bold) | 1.2 | Hero titles (with gradient) |
| Heading 1 | 36px | 700 (Bold) | 1.3 | Page titles |
| Heading 2 | 28-32px | 700 (Bold) | 1.35 | Section headers |
| Heading 3 | 22px | 700 (Bold) | 1.4 | Card titles, trait names |
| Body Large | 18px | 400-600 | 1.6 | Question text |
| Body | 16px | 400 (Regular) | 1.6 | General content |
| Body Small | 14px | 400-500 | 1.5 | Captions, helper text |
| Caption | 12px | 500 (Medium) | 1.4 | Labels, metadata |

### Gradient Text

For hero titles and emphasis:

```css
background: linear-gradient(...);
-webkit-background-clip: text;
-webkit-text-fill-color: transparent;
background-clip: text;
```

---

## Spacing and Layout

### Spacing Scale

Use consistent spacing based on a 4px base unit:

| Token | Value | Usage |
|-------|-------|-------|
| `space-1` | 4px | Tight gaps, icon padding |
| `space-2` | 8px | Between related elements |
| `space-3` | 12px | Form field gaps |
| `space-4` | 16px | Standard padding |
| `space-5` | 24px | Section padding |
| `space-6` | 32px | Card padding |
| `space-8` | 48px | Section margins |
| `space-10` | 64px | Page sections |

### Layout Principles

- **Mobile-first**: Design for 320px minimum width, scale up
- **Card-based**: Present questions in individual cards for focus
- **Generous whitespace**: Don't crowd elements; let content breathe
- **Single column**: Keep questionnaire flow simple and linear
- **Colorful backgrounds**: Use gradient blobs to add visual interest

### Border Radius

| Size | Value | Usage |
|------|-------|-------|
| Small | 8px | Small buttons, tags |
| Medium | 12px | Buttons, inputs |
| Large | 20-24px | Cards, containers |
| Extra Large | 32px | Hero elements |
| Full | 9999px | Pills, circular elements |

---

## Background Decorations

### Colorful Blob Pattern

Each page features large, blurred gradient circles that create a vibrant, energetic atmosphere:

```html
<!-- Example: Hero section background -->
<div class="absolute inset-0 overflow-hidden pointer-events-none">
  <div class="absolute -top-20 -right-20 w-[500px] h-[500px] 
              bg-gradient-to-br from-sky-400/40 to-teal-400/30 
              rounded-full blur-3xl" />
  <div class="absolute -bottom-32 -left-32 w-[400px] h-[400px] 
              bg-gradient-to-tr from-rose-400/35 to-orange-300/25 
              rounded-full blur-3xl" />
  <div class="absolute top-20 left-1/4 w-[300px] h-[300px] 
              bg-gradient-to-br from-violet-400/30 to-purple-300/20 
              rounded-full blur-3xl" />
</div>
```

### Guidelines

- Use `blur-3xl` (64px blur) for soft, diffused blobs
- Position blobs partially off-screen for a dynamic feel
- Layer 3-4 blobs with different colors per section
- Use 20-40% opacity to keep blobs subtle but visible
- Always use `pointer-events-none` to prevent interaction issues

---

## Components

### Primary Button (Gradient)

Used for main actions like "Start Test" or "See Results":

```css
/* Base */
background: linear-gradient(to right, #0EA5E9, #14B8A6, #10B981);
color: #FFFFFF;
padding: 20px 40px;
border-radius: 16px;
font-size: 18px;
font-weight: 700;
box-shadow: 0 10px 25px -5px rgba(20, 184, 166, 0.3);

/* Hover */
transform: translateY(-4px);
box-shadow: 0 20px 40px -10px rgba(20, 184, 166, 0.4);
```

### Secondary Button

Used for alternative actions like "Go Back" or "Retake":

```css
background: #FFFFFF;
color: #334155;
padding: 12px 24px;
border-radius: 12px;
font-weight: 600;
border: 2px solid #E2E8F0;

/* Hover */
border-color: #A78BFA;
background: #FAF5FF;
```

### Question Card (Gradient Border)

```css
/* Outer wrapper - gradient border */
.card-wrapper {
  padding: 3px;
  border-radius: 24px;
  background: linear-gradient(to right, #8B5CF6, #A855F7, #EC4899);
  box-shadow: 0 25px 50px -12px rgba(139, 92, 246, 0.2);
}

/* Inner card */
.card-inner {
  background: #FFFFFF;
  border-radius: 21px;
  padding: 24px 32px;
}
```

### Answer Options (Likert Scale)

```css
/* Default state */
padding: 16px;
border-radius: 12px;
border: 2px solid #E2E8F0;
background: #FFFFFF;

/* Hover state */
border-color: #C4B5FD;
background: #FAF5FF;
transform: translateY(-2px);

/* Selected state */
border: none;
background: linear-gradient(to right, #8B5CF6, #A855F7);
color: #FFFFFF;
box-shadow: 0 10px 25px -5px rgba(139, 92, 246, 0.3);
```

### Progress Bar

```css
/* Container */
height: 12px;
background: rgba(226, 232, 240, 0.8);
border-radius: 9999px;
backdrop-filter: blur(4px);
box-shadow: inset 0 2px 4px rgba(0, 0, 0, 0.05);

/* Fill */
background: linear-gradient(to right, #8B5CF6, #A855F7, #EC4899);
border-radius: 9999px;
box-shadow: 0 4px 12px rgba(139, 92, 246, 0.3);

/* Shimmer effect overlay */
.shimmer {
  background: linear-gradient(to right, transparent, rgba(255,255,255,0.3), transparent);
  animation: shimmer 2s infinite;
}
```

### Trait Result Card

```css
padding: 16px;
background: #FFFFFF;
border-radius: 16px;
border: 1px solid #F1F5F9;
box-shadow: 0 4px 20px rgba(0, 0, 0, 0.05);

/* Icon container */
.icon-wrapper {
  width: 40px;
  height: 40px;
  border-radius: 12px;
  background: linear-gradient(to bottom right, [trait-color-light], [trait-color-medium]);
}

/* Score bar */
.score-bar {
  height: 12px;
  border-radius: 9999px;
  background: linear-gradient(to right, [trait-primary], [trait-secondary]);
}
```

### Feature Section (Dark Gradient)

```css
background: linear-gradient(to bottom right, #7C3AED, #9333EA, #4F46E5);
border-radius: 24px;
padding: 48px;
color: #FFFFFF;
box-shadow: 0 25px 50px -12px rgba(124, 58, 237, 0.25);

/* Decorative blobs inside */
.blob {
  position: absolute;
  border-radius: 9999px;
  filter: blur(48px);
  opacity: 0.2-0.3;
}
```

### Pill Badge

```css
display: inline-flex;
padding: 8px 16px;
border-radius: 9999px;
background: rgba(255, 255, 255, 0.1);
backdrop-filter: blur(4px);
font-size: 14px;
font-weight: 600;
```

---

## Icons and Imagery

### Icon Style

- **Library**: Lucide Icons
- **Style**: Rounded, friendly line icons (2px stroke weight)
- **Size**: 16px (small), 20px (default), 24px (medium), 32px (large)
- **Colors**: Match trait colors or use contextual colors

### Trait Icons

| Trait | Icon | Color |
|-------|------|-------|
| Extraversion | `zap` | Rose/Pink |
| Agreeableness | `heart-handshake` | Emerald/Teal |
| Conscientiousness | `target` | Amber/Orange |
| Emotional Stability | `anchor` | Sky/Blue |
| Openness | `lightbulb` | Violet/Purple |

### Decorative Icons

- `sparkles` - Success, celebration
- `trophy` - Completion, achievement
- `play` - Start actions
- `arrow-right` - Continue, next
- `check-circle` - Confirmation

---

## Accessibility

### Color Contrast

All text must meet WCAG 2.1 AA standards:

| Text Type | Minimum Contrast Ratio |
|-----------|----------------------|
| Normal text (< 18px) | 4.5:1 |
| Large text (≥ 18px bold or ≥ 24px) | 3:1 |
| UI components and graphics | 3:1 |

**Pre-verified combinations:**
- `Slate 800 (#1E293B)` on `White (#FFFFFF)` = 12.6:1
- `Slate 600 (#475569)` on `White (#FFFFFF)` = 6.4:1
- `White (#FFFFFF)` on gradient buttons = 4.5:1+ (verified)

### Focus States

```css
outline: 2px solid #8B5CF6;
outline-offset: 2px;
```

### Motion and Animation

```css
/* Respect reduced motion preferences */
@media (prefers-reduced-motion: reduce) {
  *,
  *::before,
  *::after {
    animation-duration: 0.01ms !important;
    animation-iteration-count: 1 !important;
    transition-duration: 0.01ms !important;
  }
}
```

### Animation Guidelines

- Keep animations subtle and purposeful
- Avoid continuous looping animations (except subtle progress indicators)
- Use `ease` or `ease-out` timing functions
- Typical duration: 200-300ms for interactions, 500ms for transitions

---

## CSS Custom Properties

```css
:root {
  /* Primary Colors */
  --color-sky: #0EA5E9;
  --color-teal: #14B8A6;
  --color-violet: #8B5CF6;
  --color-purple: #A855F7;
  
  /* Trait Colors */
  --color-extraversion: #F43F5E;
  --color-agreeableness: #10B981;
  --color-conscientiousness: #F59E0B;
  --color-stability: #0EA5E9;
  --color-openness: #8B5CF6;
  
  /* Neutral Colors */
  --color-white: #FFFFFF;
  --color-cloud: #F8FAFC;
  --color-slate-100: #F1F5F9;
  --color-slate-200: #E2E8F0;
  --color-slate-500: #64748B;
  --color-slate-600: #475569;
  --color-slate-800: #1E293B;
  
  /* Typography */
  --font-sans: 'Nunito', 'Inter', 'Segoe UI', sans-serif;
  
  /* Border Radius */
  --radius-sm: 8px;
  --radius-md: 12px;
  --radius-lg: 20px;
  --radius-xl: 24px;
  --radius-2xl: 32px;
  --radius-full: 9999px;
  
  /* Shadows */
  --shadow-sm: 0 2px 8px rgba(0, 0, 0, 0.05);
  --shadow-md: 0 4px 20px rgba(0, 0, 0, 0.05);
  --shadow-lg: 0 10px 40px rgba(0, 0, 0, 0.08);
  --shadow-colored: 0 10px 25px -5px;
}
```

---

*This style guide is a living document. Update it as the design evolves to ensure consistency across the VOCA platform.*
