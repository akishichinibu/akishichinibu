import typing as t

skill_and_progress: t.List[t.Tuple[str, t.Optional[str], str]] = [
  ["python", None, ""],
  ["typescript", None, ""],
  ["cplusplus", None, ""],
  ["java", None, ""],
  ["kotlin", None, ""],
  ["go", None, ""],
  ["rust", "rust-plain", ""],
  ["html5", None, ""],
  ["css3", None, ""], 
  ["docker", None, ""],
  ["django", None, ""],
  ["flask", None, ""],
  ["react", None, ""],
  ["vuejs", None, ""],
]


def get_image_url(skill: str, fn: t.Optional[str]=None):
  return f'https://cdn.jsdelivr.net/gh/devicons/devicon/icons/{skill}/{skill + "-original" if fn is None else fn}.svg'


def generate_skill_matrix(skills: t.List[t.Tuple[str, float]], columnNums: int, size: int):
  n = len(skills)
  rowNums = int(n / columnNums) + (0 if n % columnNums == 0 else 1)

  buffer = [
    '<div>',
    '</div>',
  ]

  for r in range(rowNums):
    buffer.append('<div>')
    for c in range(columnNums):
      t = r * columnNums + c
      if t >= n:
        break
      skill, fn, status = skills[t]
      buffer.append(f'<img src="{get_image_url(skill, fn)}" height={size} width={size} />')

    buffer.append('</div>')
  
  return '\n'.join(buffer)

content = f"""
## Hi there 👋

### About Me

I'm akishichinibu, a full-stack engineer. 

I'm now living in Tokyo, Japan. 

### Skill Stack
{generate_skill_matrix(skill_and_progress, 5, 50)}

and also AWS, etc.
"""

print(content)
