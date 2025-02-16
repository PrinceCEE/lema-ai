export const postTextFormatter = (text: string, isTitle: boolean): string => {
  const wordCount = text.split(" ").length;
  if (wordCount === 1) {
    return text.slice(0, 17) + "...";
  }

  if (isTitle) {
    return text.length > 50 ? text.slice(0, 50) + "..." : text;
  }

  return text;
};
