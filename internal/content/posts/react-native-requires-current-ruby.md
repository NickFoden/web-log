Hitting issues with React Native setup or errors when running the out of the box starters? You may need to update or check your ruby version. This caught me up the other day on a new machine.

1. Install Ruby:

   ```
   brew install ruby
   ```

2. Get the prefix path for Ruby:

   ```
   brew --prefix ruby
   ```

3. Add Ruby to your PATH:

   ```
   echo 'export PATH="$(brew --prefix ruby)/bin:$PATH"' >> ~/.zshrc
   ```

4. Source your updated .zshrc file:

   ```
   source ~/.zshrc
   ```

5. Verify Ruby installation:

   ```
   ruby -v
   ```

6. Install CocoaPods:

   ```
   gem install cocoapods
   ```
