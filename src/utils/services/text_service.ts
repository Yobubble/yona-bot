export class TextService {
  static postProcessing(text: string): string {
    const textFragments = text.split("<br>").join("");
    const processedText = textFragments.replace(/<i>(.*?)<\/i>/g, "**$1**");
    return processedText;
  }
}
